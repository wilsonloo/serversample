<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Login Test</title>
		<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
		<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
		<script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
		<script src="//cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	</head>
	<body>
		<div class="container">
			<h1 class="page-header">
			认证服务器—用户登录&nbsp;
			<small>默认的登录账号(用户名：admin,密码：admin)</small>
			</h1>
			{{if .ErrorMessage}}
			<div class="alert alert-danger" role="alert">{{.ErrorMessage}}</div>
			{{end}}
			<form action="/login?code=123" method="post">
				<div class="form-group">
					<label for="UserName">用户名</label>
					<input type="text" class="form-control" name="UserName" placeholder="请输入用户名">
				</div>
				<div class="form-group">
					<label for="Password">密码</label>
					<input type="password" class="form-control" name="Password" placeholder="请输入密码">
				</div>
				<button type="submit" class="btn btn-primary">登录</button>
			</form>
		</div>
	</body>
</html>