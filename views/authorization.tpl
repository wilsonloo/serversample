<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>应用授权</title>
		<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
		<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
		<script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
		<script src="//cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	</head>
	<body>
		<div class ="container">
			<div class="jumbotron">
				<h1>{{.ClientName}}</h1>
				<p>授权提示：将允许该应用获取您的个人信息!</p>
				<p>
					<form action="/authorize" method="post">
						<button type="submit" class="btn btn-danger btn-lg">授权</button>&nbsp;
						<button type="button" class="btn btn-default btn-lg" onclick="window.close();">关闭</button>
					</form>
				</a>
			</p>
		</div>
	</div>
</body>
</html>