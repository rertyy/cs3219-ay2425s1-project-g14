// contains the database-related functions for the questions API.
package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *QuestionDB) GetAllQuestions(logger *Logger) ([]Question, error) {
	questionCursor, err := db.questions.Find(context.Background(), bson.D{})

	if err != nil {
		logger.Log.Error("Error retrieving questions: ", err.Error())
		return nil, err
	}

	var questions []Question

	
	if err = questionCursor.All(context.Background(), &questions); err != nil {
		logger.Log.Error("Error decoding questions: ", err.Error())
		return nil, err
	}

	return questions, nil
}