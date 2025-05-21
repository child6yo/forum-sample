package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/child6yo/forum-sample"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateThread(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unable to make mock db: %s", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewThreadsDatabase(sqlxDB)

	type args struct {
		postId int
		thread forum.Threads
	}
	type mockBehavior func(args args, id int)

	testCases := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	} {
		{
			name: "Ok",
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO threads").
					WithArgs(
						args.thread.UserId, args.thread.Content, args.thread.AnswerAt, 
						args.thread.CrTime, args.thread.Update, args.thread.UpdTime).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO post_threads").WithArgs(args.postId, id).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			input: args{
				postId: 1,
				thread: forum.Threads{
					UserId: 1,
					Content: "content",
					AnswerAt: 0,
					CrTime: time.Time{},
					Update: false,
					UpdTime: time.Time{},
				},
			},
			want: 1,
			wantErr: false,
		},
		{
			name: "Empty Fields",
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id).RowError(0, errors.New("insert error"))
				mock.ExpectQuery("INSERT INTO threads").
				WithArgs(
					args.thread.UserId, args.thread.Content, args.thread.AnswerAt, 
					args.thread.CrTime, args.thread.Update, args.thread.UpdTime).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			input: args{
				postId: 1,
				thread: forum.Threads{
					UserId: 1,
					Content: "",
					AnswerAt: 0,
					CrTime: time.Time{},
					Update: false,
					UpdTime: time.Time{},
				},
			},
			wantErr: true,
		},
		{
			name: "Failed 2nd Insert",
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO threads").
				WithArgs(
					args.thread.UserId, args.thread.Content, args.thread.AnswerAt, 
					args.thread.CrTime, args.thread.Update, args.thread.UpdTime).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO post_threads").WithArgs(args.postId, id).
					WillReturnError(errors.New("insert error"))

				mock.ExpectRollback()
			},
			input: args{
				postId: 1,
				thread: forum.Threads{
					UserId: 1,
					Content: "content",
					AnswerAt: 0,
					CrTime: time.Time{},
					Update: false,
					UpdTime: time.Time{},
				},
			},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		test.mock(test.input, test.want)

		got, err := r.CreateThread(test.input.postId, test.input.thread)
		if test.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	}
}
