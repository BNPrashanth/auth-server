package helpers

// IndexPage Constant
const IndexPage = `
<html>
    <head>
        <title>Test JWT</title>
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/css/bootstrap.min.css">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.0/jquery.min.js"></script>
		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/js/bootstrap.min.js"></script>
		<script type="text/javascript">

			var accessToken = "";

			signin = () => {
				console.log("Signin Clicked..");
				$.ajax({
					url: "http://localhost:9090/signin",
					type: 'post',
					data: JSON.stringify({
							username: "user2",
							password: "password2"
						}),
					success: function(data) {
						result = JSON.parse(data);
						console.log(result);
						accessToken = result.Value;
						console.log(accessToken);
					}
				});
			}

			welcome = () => {
				console.log("Welcome Clicked..");
				$.ajax({
					url: "http://localhost:9090/welcome",
					type: 'post',
					headers: {
						Authorization: "bearer " + accessToken
					},
					success: function(data, status){
						result = JSON.parse(data);
						console.log(result);
					}
				});
			}

			refresh = () => {
				console.log("Refresh Clicked..");
				$.ajax({
					url: "http://localhost:9090/refresh",
					type: 'post',
					headers: {
						Authorization: "bearer " + accessToken
					},
					success: function(data, status){
						result = JSON.parse(data);
						console.log(result);
						accessToken = result.Value;
						console.log(accessToken);
					}
				});
			}
		</script>
    </head>
    <body>
        <h2>Testing How JWT Works in GoLang</h2>
        <p>
            <ul>
                <li onclick="signin()" style="cursor: pointer;">
					Signin
				</li>
				<li onclick="welcome()" style="cursor: pointer;">
					Welcome
				</li>
				<li onclick="refresh()" style="cursor: pointer;">
					Refresh
				</li>
            </ul>
        </p>
    </body>
</html>
`
