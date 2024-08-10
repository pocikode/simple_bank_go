package mail

import (
	"github.com/stretchr/testify/require"
	"pocikode/simple-bank/util"
	"testing"
)

func TestSendEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewEmailSender(
		config.SMTPSenderName,
		config.SMTPSenderAddress,
		config.SMTPHost,
		config.SMTPPort,
		config.SMTPUsername,
		config.SMTPPassword,
	)

	subject := "A test email"
	content := `
	<h1>Hello World</h1>
	<p>This is a test message.</p>
	`

	to := []string{"agus@pocikode.dev"}
	attachFiles := []string{"../README.md"}

	err = sender.Send(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
