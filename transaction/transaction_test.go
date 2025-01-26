package transaction

import (
	"testing"
	"ticketing-system/config"
	"ticketing-system/entity"

	"github.com/stretchr/testify/assert"
)

func TestProcessTransaction(t *testing.T) {
	// Setup test database
	config.SetupTestDB()
	defer config.TeardownTestDB()

	// Gunakan TestDB untuk operasi database
	db := config.TestDB

	// Create a mock event
	event := entity.Event{
		Name:     "Concert A",
		Price:    100.0,
		Capacity: 50,
	}
	db.Create(&event)

	// Test case: Successful transaction
	t.Run("Success", func(t *testing.T) {
		ticket, err := ProcessTransaction(1, event.ID, 10)
		assert.NoError(t, err)
		assert.NotNil(t, ticket)
		assert.Equal(t, 10, ticket.Quantity)
	})
}
