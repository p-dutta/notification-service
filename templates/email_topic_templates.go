package subscribers

var SignInHtmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>OTP Email Template</title>
</head>
<body style="font-family: Arial, sans-serif;">
	<div style="max-width: 600px; margin: 0 auto;">
		<h2>Your OTP for verification:</h2>
		<h1 style="font-size: 36px; background-color: #f5f5f5; padding: 10px 20px; display: inline-block;">
			<!-- Place OTP code here -->
			OTP_HERE
		</h1>
		<p>Please use the above OTP code to verify your account.</p>
		<p>If you did not request this OTP, please ignore this message.</p>
	</div>
</body>
</html>`
