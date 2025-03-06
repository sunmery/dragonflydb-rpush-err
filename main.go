package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/redis-adapter/v3"
)

func main() {
	// Direct Initialization:
	// Initialize a Redis adapter and use it in a Casbin enforcer:
	// a, err1 := redisadapter.NewAdapterWithPassword("tcp", "localhost:6379", "your_secure_password") // Your Redis network and address.
	a, err1 := redisadapter.NewAdapter("tcp", "localhost:6379") // Your Redis network and address.
	if err1 != nil {
		panic(fmt.Errorf("failed to initialize redis adapter: %v", err1))
	}

	// Use the following if Redis has password like "123"
	// a, err := redisadapter.NewAdapterWithPassword("tcp", "127.0.0.1:6379", "123")

	// Use the following if you use Redis with a specific user
	// a, err := redisadapter.NewAdapterWithUser("tcp", "127.0.0.1:6379", "username", "password")

	// Use the following if you use Redis connections pool
	// pool := &redis.Pool{}
	// a, err := redisadapter.NewAdapterWithPool(pool)

	// Initialization with different user options:
	// Use the following if you use Redis with passowrd like "123":
	// a, err := redisadapter.NewAdapterWithOption(redisadapter.WithNetwork("tcp"), redisadapter.WithAddress("127.0.0.1:6379"), redisadapter.WithPassword("123"))

	// Use the following if you use Redis with username, password, and TLS option:
	// var clientTLSConfig tls.Config
	// ...
	// a, err := redisadapter.NewAdapterWithOption(redisadapter.WithNetwork("tcp"), redisadapter.WithAddress("127.0.0.1:6379"), redisadapter.WithUsername("testAccount"), redisadapter.WithPassword("123456"), redisadapter.WithTls(&clientTLSConfig))

	e, err2 := casbin.NewEnforcer("./rbac_model.conf", a)
	if err2 != nil {
		panic(fmt.Errorf("failed to initialize enforcer: %v", err2))
	}
	fmt.Printf("e: %+v\n", e)

	// Load the policy from DB.
	err3 := e.LoadPolicy()
	if err3 != nil {
		panic(fmt.Errorf("failed to load policy: %v", err3))
	}

	// Check the permission.
	enforce, err4 := e.Enforce("alice", "data1", "read")
	if err4 != nil {
		panic(fmt.Errorf("failed to enforce: %v", err4))
	}

	fmt.Println(enforce)
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	err5 := e.SavePolicy()
	if err5 != nil {
		panic(fmt.Errorf("failed to save policy: %v", err5))
	}
}

