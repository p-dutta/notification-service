package subscribers

var EmailAdminHtmlTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <title>OTP Email Template</title>
	</head>
	<body style="font-family: Arial, sans-serif;">
	    <div style="max-width: 600px; margin: 0 auto;">
	        <h2>Email Admin :</h2>
	        
	    </div>
	</body>
	</html>`

var EmailVerificationHtmlTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <title>OTP Email Template</title>
	</head>
	<body style="font-family: Arial, sans-serif;">
	    <div style="max-width: 600px; margin: 0 auto;">
	        <h2>Email verification:</h2>
	       
	    </div>
	</body>
	</html>`

var ForgetPasswordHtmlTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Forgot Password</title>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				background-color: #f4f4f4;
				margin: 0;
				padding: 0;
				text-align: center;
			}
	
			.container {
				max-width: 600px;
				margin: 50px auto;
				background-color: #ffffff;
				padding: 20px;
				border-radius: 8px;
				box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			}
	
			h2 {
				color: #333;
			}
	
			p {
				color: #666;
				line-height: 1.6;
			}
	
			.button {
				display: inline-block;
				padding: 10px 20px;
				margin: 20px 0;
				font-size: 16px;
				text-decoration: none;
				background-color: #3498db;
				color: #fff;
				border-radius: 5px;
			}
	
			.footer {
				margin-top: 20px;
				color: #888;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Forgot Your Password?</h2>
			<p>No worries! It happens to the best of us. To reset your password, simply click the button below:</p>
			<a class="button" href="YOUR_RESET_URL">Reset Password</a>
			<p>If you did not request a password reset, please ignore this email.</p>
			<div class="footer">
				<p>Best regards,<br>Your Company Name</p>
			</div>
		</div>
	</body>
	</html>`

var ResetPasswordHtmlTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>Password Reset</title>
	  <style>
		body {
		  font-family: Arial, sans-serif;
		  background-color: #f4f4f4;
		  margin: 0;
		  padding: 0;
		  display: flex;
		  align-items: center;
		  justify-content: center;
		  height: 100vh;
		}
	
		.reset-container {
		  background-color: #fff;
		  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		  border-radius: 8px;
		  padding: 20px;
		  width: 300px;
		  text-align: center;
		}
	
		h2 {
		  color: #333;
		}
	
		form {
		  display: flex;
		  flex-direction: column;
		  margin-top: 20px;
		}
	
		label {
		  margin-bottom: 8px;
		  color: #555;
		}
	
		input {
		  padding: 10px;
		  margin-bottom: 16px;
		  border: 1px solid #ccc;
		  border-radius: 4px;
		  font-size: 14px;
		}
	
		button {
		  background-color: #007BFF;
		  color: #fff;
		  padding: 12px;
		  border: none;
		  border-radius: 4px;
		  cursor: pointer;
		  font-size: 16px;
		}
	
		button:hover {
		  background-color: #0056b3;
		}
	  </style>
	</head>
	<body>
	
	<div class="reset-container">
	  <h2>Password Reset</h2>
	  <form action="reset_password.php" method="post">
		<label for="email">Email:</label>
		<input type="email" id="email" name="email" required>
	
		<label for="newPassword">New Password:</label>
		<input type="password" id="newPassword" name="newPassword" required>
	
		<button type="submit">Reset Password</button>
	  </form>
	</div>
	
	</body>
	</html>`

var AdminSignInHtmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>OTP Admin Email Template</title>
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

var SmsAdminHtmlTemplate = `<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <title>OTP Email Template</title>
	</head>
	<body style="font-family: Arial, sans-serif;">
	    <div style="max-width: 600px; margin: 0 auto;">
	        <h2>Sms Admin:</h2>
	        
	    </div>
	</body>
	</html>`
