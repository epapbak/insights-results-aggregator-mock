/*
Copyright © 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

// Generated documentation is available at:
// https://godoc.org/github.com/RedHatInsights/insights-results-aggregator-mock/types
//
// Documentation in literate-programming-style is available at:
// https://redhatinsights.github.io/insights-results-aggregator-mock/packages/types/errors.html

import (
	"errors"
	"fmt"
)

// ErrOldReport is an error returned if a more recent already
// exists on the storage while attempting to write a report for a cluster.
var ErrOldReport = errors.New("More recent report already exists in storage")

// ItemNotFoundError shows that item with id ItemID wasn't found in the storage
type ItemNotFoundError struct {
	ItemID interface{}
}

// Error returns error string
func (e *ItemNotFoundError) Error() string {
	return fmt.Sprintf("Item with ID %+v was not found in the storage", e.ItemID)
}

// TableNotFoundError table not found error
type TableNotFoundError struct {
	tableName string
}

// Error returns error string
func (err *TableNotFoundError) Error() string {
	return fmt.Sprintf("no such table: %v", err.tableName)
}

// TableAlreadyExistsError represents table already exists error
type TableAlreadyExistsError struct {
	tableName string
}

// Error returns error string
func (err *TableAlreadyExistsError) Error() string {
	return fmt.Sprintf("table %v already exists", err.tableName)
}

// ForeignKeyError something violates foreign key error
// tableName and foreignKeyName can be empty for DBs not supporting it (SQLite)
type ForeignKeyError struct {
	TableName      string
	ForeignKeyName string

	// Details can reveal you information about specific item violating fk
	Details string
}

// Error returns error string
func (err *ForeignKeyError) Error() string {
	return fmt.Sprintf(
		`operation violates foreign key "%v" on table "%v"`, err.ForeignKeyName, err.TableName,
	)
}

/*
// ConvertDBError converts sql errors to those defined in this package
func ConvertDBError(err error, itemID interface{}) error {
	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		if itemIDArray, ok := itemID.([]interface{}); ok {
			var strArray []string
			for _, item := range itemIDArray {
				strArray = append(strArray, fmt.Sprint(item))
			}

			itemID = strings.Join(strArray, "/")
		}

		return &ItemNotFoundError{ItemID: itemID}
	}

	err = convertPostgresError(err)
	err = convertSQLiteError(err)

	return err
}
*/
