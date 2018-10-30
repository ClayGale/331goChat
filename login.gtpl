<!DOCTYPE html>
<html>
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" type="text/css" href="login.css">
	<link href='http://fonts.googleapis.com/css?family=Just+Another+Hand' rel='stylesheet' type='text/css'>

</head>
<body>
	<h2><center> Lets Talk All Together</center></h2>

	<div id="container">
		<form action="/login">
			<label for="username">Create a Username:</label>
			<input type="text" id="username" name="name" maxlength="20" required>
			<p style="padding-left: 20px">Choose a User Colour:</p>

			<select id="soflow" name="colour">
				<option value="0" selected disabled>Select your User Colour</option>
				<option value="red">Red</option> <!--#FF0000-->
				<option value="orange">Orange</option><!--#FFA500 -->
				<option value="yellow">Yellow</option><!--#FFFF00 -->
				<option value="green">Green</option><!-- #008000-->
				<option value="blue">Blue</option><!-- #0000FF-->
				<option value="indigo">Indigo</option><!-- #4b0082-->
				<option value="violet">Violet</option><!-- #9400D3-->
			</select>

			<div class=button>
				<input type="submit" class="submit" value="Click to Chat">
			</div>

		</form>
	</div>

</body>
</html>
