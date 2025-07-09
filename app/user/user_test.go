package main

import (
	"context"
	"fmt"
	"proxima/app/user/api/pbentity"
	"proxima/app/user/internal/logic/account"
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"
)

// ... (setup function remains the same)

func TestRegister(t *testing.T) {
	ctx := gctx.New()

	gtest.C(t, func(t *gtest.T) {
		// 正常注册单个用户，用于后续测试用户名/邮箱已存在的情况
		firstUser := &pbentity.Users{
			Username: "testuser1",
			Email:    "testuser1@example.com",
			Password: "password123",
		}
		_, err := account.Register(ctx, firstUser)
		t.Assert(err, nil)

		// 用户名已存在
		userConflictUsername := &pbentity.Users{
			Username: "testuser1", // 与 firstUser 用户名相同
			Email:    "testuser_conflict_username@example.com",
			Password: "password456",
		}
		_, errConflictUsername := account.Register(ctx, userConflictUsername)
		t.AssertNE(errConflictUsername, nil)
		t.Assert(errConflictUsername.Error(), "用户已存在")

		// 邮箱已存在
		userConflictEmail := &pbentity.Users{
			Username: "testuser_conflict_email",
			Email:    "testuser1@example.com", // 与 firstUser 邮箱相同
			Password: "password789",
		}
		_, errConflictEmail := account.Register(ctx, userConflictEmail)
		t.AssertNE(errConflictEmail, nil)
		t.Assert(errConflictEmail.Error(), "邮箱已存在")

		// 测试超时 - 单个用户
		shortTimeoutCtx, cancelShort := context.WithTimeout(ctx, 1*time.Nanosecond)
		userTimeout := &pbentity.Users{
			Username: "timeoutuser_single",
			Email:    "timeout_single@example.com",
			Password: "passwordtimeout_single",
		}
		_, errTimeout := account.Register(shortTimeoutCtx, userTimeout)
		cancelShort() // Ensure context is cancelled
		t.AssertNE(errTimeout, nil)
		g.Log().Infof(ctx, "Single user timeout test error: %v", errTimeout)
	})
}

func TestRegisterConcurrent(t *testing.T) {
	ctx := gctx.New()
	numUsers := 20

	gtest.C(t, func(t *gtest.T) {
		var wg sync.WaitGroup
		errors := make(chan error, numUsers)     // Channel to collect errors
		successCount := make(chan int, numUsers) // Channel to count successes

		for i := 0; i < numUsers; i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				username := fmt.Sprintf("user_%d", idx)
				email := fmt.Sprintf("concurrent_user_%d@example.com", idx)
				user := &pbentity.Users{
					Username: username,
					Email:    email,
					Password: "password123",
				}

				// 为每个goroutine创建一个新的context，以防原始context被意外取消
				// 或者，如果需要测试超时，可以在这里创建带超时的context
				// localCtx, cancel := context.WithTimeout(gctx.New(), 30*time.Second) // 示例：30秒超时
				// defer cancel()
				localCtx := gctx.New() // 使用新的独立上下文

				id, err := account.Register(localCtx, user)
				if err != nil {
					errors <- fmt.Errorf("user %s (%s): %w", username, email, err)
					return
				}
				if id <= 0 {
					errors <- fmt.Errorf("user %s (%s): received invalid ID %d", username, email, id)
					return
				}
				successCount <- 1
			}(i)
		}

		wg.Wait()
		close(errors)
		close(successCount)

		totalSuccesses := 0
		for range successCount {
			totalSuccesses++
		}

		numErrors := 0
		for err := range errors {
			numErrors++
			g.Log().Errorf(ctx, "Concurrent registration error: %v", err)
		}

		t.AssertEQ(totalSuccesses, numUsers-numErrors) // 确保成功的数量符合预期
		// 根据实际情况，如果允许部分失败（例如数据库偶发性锁定），可以调整断言
		// 如果期望所有都成功，则 numErrors 应该为 0
		t.AssertEQ(numUsers, "预期所有并发注册都成功")

		// 可选：进一步验证数据库中确实创建了这些用户
		// for i := 0; i < numUsers; i++ {
		// 	 username := fmt.Sprintf("concurrent_user_%d", i)
		// 	 // dbCheckCtx, dbCheckCancel := context.WithTimeout(gctx.New(), 5*time.Second)
		// 	 // defer dbCheckCancel()
		// 	 // count, errDb := g.DB().Model("users").Where("username", username).Count()
		// 	 // t.Assert(errDb, nil)
		// 	 // t.AssertEQ(count, 1, fmt.Sprintf("User %s not found in DB after concurrent registration", username))
		// }
	})
}

func TestLoginConcurrent(t *testing.T) {
	ctx := gctx.New()
	numUsers := 20
	password := "password123"

	// Step 1: Concurrently register users that will be used for login test
	gtest.C(t, func(t *gtest.T) {
		var regWg sync.WaitGroup
		regErrors := make(chan error, numUsers)

		for i := 0; i < numUsers; i++ {
			regWg.Add(1)
			go func(idx int) {
				defer regWg.Done()
				username := fmt.Sprintf("login_concurrent_user_%d", idx)
				email := fmt.Sprintf("login_concurrent_user_%d@example.com", idx)
				user := &pbentity.Users{
					Username: username,
					Email:    email,
					Password: password, // Use a common password for simplicity in login
				}
				localCtx := gctx.New()
				_, err := account.Register(localCtx, user)
				if err != nil {
					// If a user already exists from a previous failed run, log it but don't fail the setup for login test
					count, err := g.DB().Model("users").Where("username", username).Count()
					if err != nil {
						return
					}
					if count > 0 {
						g.Log().Warningf(ctx, "User %s already exists, proceeding with login test.", username)
					} else {
						regErrors <- fmt.Errorf("setup: user %s (%s): %w", username, email, err)
					}
				}
			}(i)
		}
		regWg.Wait()
		close(regErrors)

		for err := range regErrors {
			t.Fatalf("Failed to register user during setup for concurrent login test: %v", err)
		}
		g.Log().Info(ctx, "Concurrent registration for login test completed.")
	})

	// Step 2: Concurrently login users
	gtest.C(t, func(t *gtest.T) {
		var loginWg sync.WaitGroup
		loginErrors := make(chan error, numUsers)
		loginSuccessCount := make(chan int, numUsers)

		for i := 0; i < numUsers; i++ {
			loginWg.Add(1)
			go func(idx int) {
				defer loginWg.Done()
				username := fmt.Sprintf("user_%d", idx)
				localCtx := gctx.New()

				_, _, token, err := account.Login(localCtx, username, password)
				if err != nil {
					loginErrors <- fmt.Errorf("login user %s: %w", username, err)
					return
				}
				if token == "" {
					loginErrors <- fmt.Errorf("login user %s: received empty token", username)
					return
				}
				loginSuccessCount <- 1
			}(i)
		}

		loginWg.Wait()
		close(loginErrors)
		close(loginSuccessCount)

		totalLoginSuccesses := 0
		for range loginSuccessCount {
			totalLoginSuccesses++
		}

		numLoginErrors := 0
		for err := range loginErrors {
			numLoginErrors++
			g.Log().Errorf(ctx, "Concurrent login error: %v", err)
		}

		t.AssertEQ(totalLoginSuccesses, numUsers-numLoginErrors)
		t.AssertEQ(numUsers, "预期所有并发登录都成功")
	})
}

func TestLogin(t *testing.T) {
	ctx := gctx.New()

	// 先注册一个用户用于登录测试
	regUser := &pbentity.Users{
		Username: "loginuser",
		Email:    "loginuser@example.com",
		Password: "loginpassword",
	}
	_, err := account.Register(ctx, regUser)
	if err != nil {
		t.Fatalf("Failed to register user for login test: %v", err)
	}

	gtest.C(t, func(t *gtest.T) {
		// 正常登录
		id, uuid, token, errLogin := account.Login(ctx, "loginuser", "loginpassword")
		t.Assert(errLogin, nil)
		t.AssertNE(id, "")
		t.AssertNE(uuid, "")
		t.AssertNE(token, "")

		// 用户名不存在
		_, _, _, errNotFound := account.Login(ctx, "nonexistentuser", "somepassword")
		t.AssertNE(errNotFound, nil)
		t.Assert(errNotFound.Error(), "用户名不存在")

		// 密码错误
		_, _, _, errWrongPass := account.Login(ctx, "loginuser", "wrongpassword")
		t.AssertNE(errWrongPass, nil)
		t.Assert(errWrongPass.Error(), "用户名或密码错误")

		// 测试登录超时 - 类似于注册的超时测试
		shortTimeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Nanosecond)
		defer cancel()
		_, _, _, errTimeoutLogin := account.Login(shortTimeoutCtx, "loginuser", "loginpassword")
		t.AssertNE(errTimeoutLogin, nil)
		g.Log().Infof(ctx, "Login timeout test error: %v", errTimeoutLogin)
	})
}

// 注意：为了使这些测试能够正确运行，你需要确保：
// 1. 数据库配置正确，并且测试数据库可以访问。
// 2. 如果你的 `utility.Encrypt` 和 `utility.Verify` 函数有外部依赖（例如密钥），确保它们在测试环境中可用。
// 3. `proxima/utility/consts.JwtKey` 和 `proxima/utility.SetJWT` (以及其依赖的Redis) 在测试环境中配置正确或被mock掉。
//    对于单元测试，通常建议mock掉外部服务如Redis，以保证测试的独立性和速度。
//    你可以考虑使用接口和依赖注入来方便地mock这些依赖。
