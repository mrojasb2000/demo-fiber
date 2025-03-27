package noteHandler

import (
	"example.org/database"
	"example.org/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	// find all notes in the database
	db.Find(&notes)

	// if no note is present return an error
	if len(notes) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Notes not found", "data": nil})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Notes found", "data": notes})
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// Store the body in the note and return error if encountered
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't parse the request", "data": err})
	}

	// Add a uuid to the note
	note.ID = uuid.New()

	// Create the Note and return error if encountered
	err := db.Create(&note).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create the note", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Note created", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Note not found", "data": nil})
	}

	// Return the note with the Id
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Note found", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"text"`
	}
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Note not found", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't parse the request", "data": err})
	}

	// Edit the note
	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	// Save the Changes
	db.Save(&note)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Note updated", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Note not found", "data": nil})
	}

	// Delete the note and return error if encountered
	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": err})
	}

	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Note deleted"})
}
