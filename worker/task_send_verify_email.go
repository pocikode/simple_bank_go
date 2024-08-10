package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	db "pocikode/simple-bank/db/sqlc"
	"pocikode/simple-bank/util"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("%s: failed to marshal payload: %w", TaskSendVerifyEmail, err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("%s: failed to enqueue task: %w", TaskSendVerifyEmail, err)
	}

	log.Info().Str("type", task.Type()).
		Bytes("payload", jsonPayload).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("enqueued task")

	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("%s: failed to unmarshal payload: %w", TaskSendVerifyEmail, asynq.SkipRetry)
	}

	user, err := processor.store.GetUser(ctx, payload.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: user not found: %w", TaskSendVerifyEmail, asynq.SkipRetry)
		}

		return fmt.Errorf("%s: failed to get user: %w", TaskSendVerifyEmail, err)
	}

	verifyEmail, err := processor.store.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: util.RandomString(32),
	})
	if err != nil {
		return fmt.Errorf("%s: failed to create verify email: %w", TaskSendVerifyEmail, err)
	}

	subject := "Welcome to Simple Bank"
	verifyUrl := fmt.Sprintf("http://%s/v1/verify_email?email_id=%d&secret_code=%s", processor.config.HTTPServerAddress, verifyEmail.ID, verifyEmail.SecretCode) // FIXME: dummy url
	content := fmt.Sprintf(`Hello %s, <br/>
	Thank you for registering with us! <br/>
	Please <a href="%s">click here</a> to verify your email address. <br/>
	`, user.Fullname, verifyUrl)
	to := []string{user.Email}

	err = processor.mailer.Send(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("%s: failed to send verify email: %w", TaskSendVerifyEmail, err)
	}

	log.Info().Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("email", user.Email).
		Msg("processed task")

	return nil
}
