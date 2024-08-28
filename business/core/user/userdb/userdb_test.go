package userdb_test

import (
	"context"
	"net/mail"
	"os"
	"testing"

	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/core/user/userdb"
	"github.com/St3plox/Blogchain/foundation/web/testutil"
	"github.com/rs/zerolog"
)

var testEnv *testutil.TestEnv
var logger zerolog.Logger

func TestMain(m *testing.M) {
	ctx := context.Background()

	var err error
	testEnv, err = testutil.SetupMongoDBContainer(ctx)
	if err != nil {
		panic(err)
	}

	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	code := m.Run()

	if err := testEnv.Teardown(ctx); err != nil {
		os.Exit(1)
	}

	os.Exit(code)
}

// TestStore_Create tests the Create method using the MongoDB container
func TestStore_Create(t *testing.T) {
	ctx := context.Background()

	store := userdb.NewStore(&logger, testEnv.MongoClient)

	// Example user
	email, _ := mail.ParseAddress("test@example.com")
	usr := user.User{
		Name:  "Test User",
		Email: email.String(),
	}

	// Create user
	createdUser, err := store.Create(ctx, usr)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	if createdUser.Email != usr.Email {
		t.Errorf("Expected email %v, got %v", usr.Email, createdUser.Email)
	}
}

// TestStore_Update tests the Update method using the MongoDB container
func TestStore_Update(t *testing.T) {
	ctx := context.Background()

	store := userdb.NewStore(&logger, testEnv.MongoClient)

	// Create a user first
	email, _ := mail.ParseAddress("update_test@example.com")
	usr := user.User{
		Name:  "Update Test User",
		Email: email.String(),
	}

	createdUser, err := store.Create(ctx, usr)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Update the user's name
	createdUser.Name = "Updated User Name"
	if err := store.Update(ctx, createdUser); err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	// Verify the update
	updatedUser, err := store.QueryByID(ctx, createdUser.ID.Hex())
	if err != nil {
		t.Fatalf("Failed to query user by ID: %v", err)
	}

	if updatedUser.Name != "Updated User Name" {
		t.Errorf("Expected updated name to be %v, got %v", "Updated User Name", updatedUser.Name)
	}
}

// TestStore_Delete tests the Delete method using the MongoDB container
func TestStore_Delete(t *testing.T) {
	ctx := context.Background()

	store := userdb.NewStore(&logger, testEnv.MongoClient)

	// Create a user first
	email, _ := mail.ParseAddress("delete_test@example.com")
	usr := user.User{
		Name:  "Delete Test User",
		Email: email.String(),
	}

	createdUser, err := store.Create(ctx, usr)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Delete the user
	if err := store.Delete(ctx, createdUser); err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Verify the deletion
	_, err = store.QueryByID(ctx, createdUser.ID.Hex())
	if err == nil {
		t.Errorf("Expected error when querying deleted user, but got none")
	}
}

// TestStore_QueryByID tests the QueryByID method using the MongoDB container
func TestStore_QueryByID(t *testing.T) {
	ctx := context.Background()

	store := userdb.NewStore(&logger, testEnv.MongoClient)

	// Create a user first
	email, _ := mail.ParseAddress("query_by_id_test@example.com")
	usr := user.User{
		Name:  "QueryByID Test User",
		Email: email.String(),
	}

	createdUser, err := store.Create(ctx, usr)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Query the user by ID
	queriedUser, err := store.QueryByID(ctx, createdUser.ID.Hex())
	if err != nil {
		t.Fatalf("Failed to query user by ID: %v", err)
	}

	if queriedUser.Email != createdUser.Email {
		t.Errorf("Expected email %v, got %v", createdUser.Email, queriedUser.Email)
	}
}

// // TestStore_QueryByIDs tests the QueryByIDs method using the MongoDB container
// func TestStore_QueryByIDs(t *testing.T) {
// 	ctx := context.Background()

// 	// Connect to MongoDB instance
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
// 	if err != nil {
// 		t.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}
// 	defer client.Disconnect(ctx)

// 	// Setup logger and store
// 	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
// 	store := userdb.NewStore(&logger, client)

// 	// Create multiple users
// 	users := []user.User{
// 		{Name: "User 1", Email: "user1@example.com"},
// 		{Name: "User 2", Email: "user2@example.com"},
// 	}

// 	var userIDs []string
// 	for _, u := range users {
// 		createdUser, err := store.Create(ctx, u)
// 		if err != nil {
// 			t.Fatalf("Failed to create user: %v", err)
// 		}
// 		userIDs = append(userIDs, createdUser.ID.Hex())
// 	}

// 	time.Sleep(1000 * time.Millisecond)

// 	// Query the users by their IDs
// 	queriedUsers, err := store.QueryByIDs(ctx, userIDs)
// 	if err != nil {
// 		t.Fatalf("Failed to query users by IDs: %v", err)
// 	}

// 	if len(queriedUsers) != len(users) {
// 		t.Errorf("Expected %d users, got %d", len(users), len(queriedUsers))
// 	}
// }

// TestStore_QueryByEmail tests the QueryByEmail method using the MongoDB container
// func TestStore_QueryByEmail(t *testing.T) {
// 	ctx := context.Background()

// 	// Connect to MongoDB instance
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
// 	if err != nil {
// 		t.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}
// 	defer client.Disconnect(ctx)

// 	// Setup logger and store
// 	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
// 	store := userdb.NewStore(&logger, client)

// 	// Create a user
// 	email, _ := mail.ParseAddress("query_by_email_test@example.com")
// 	usr := user.User{
// 		Name:  "QueryByEmail Test User",
// 		Email: email.String(),
// 	}

// 	createdUser, err := store.Create(ctx, usr)
// 	if err != nil {
// 		t.Fatalf("Failed to create user: %v", err)
// 	}
// 	t.Logf("Created user with email: %s", createdUser.Email) // Log the email

// 	time.Sleep(100 * time.Millisecond)

// 	// Query the user by email
// 	queriedUser, err := store.QueryByEmail(ctx, *email)
// 	if err != nil {
// 		t.Fatalf("Failed to query user by email: %v", err)
// 	}

// 	if queriedUser.Email != createdUser.Email {
// 		t.Errorf("Expected email %v, got %v", createdUser.Email, queriedUser.Email)
// 	}
// }
