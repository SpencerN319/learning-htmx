package server 

const (
	Tmpl1 = `
	{{- define "list-}}
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<script src="https://unpkg.com/htmx.org@1.9.6" integrity="sha384-FhXw7b6AlE/jyjlZH5iHa/tTe9EpJ1Y55RjcgPbjeWMskSxZt1v9qkxLJWNJaGni" crossorigin="anonymous"></script>

			<title>{{.Title}}</title>
		</head>
		<body>
			{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
		</body>
	</html>`

	Tmpl2 = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<script src="https://unpkg.com/htmx.org@1.9.6" integrity="sha384-FhXw7b6AlE/jyjlZH5iHa/tTe9EpJ1Y55RjcgPbjeWMskSxZt1v9qkxLJWNJaGni" crossorigin="anonymous"></script>

			<title>{{.Title}}</title>
		</head>
		<body>
		<div hx-target="this" hx-swap="outerHTML">
			<div><label>First Name</label>: {{.FirstName}}</div>
			<div><label>Last Name</label>: {{.LastName}}</div>
			<div><label>Email</label>: {{.Email}}</div>
			<button hx-get="/contact/edit" class="btn btn-primary">
				Click To Edit
			</button>
		</div>
		</body>
	</html>`

	Tmpl3 = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<script src="https://unpkg.com/htmx.org@1.9.6" integrity="sha384-FhXw7b6AlE/jyjlZH5iHa/tTe9EpJ1Y55RjcgPbjeWMskSxZt1v9qkxLJWNJaGni" crossorigin="anonymous"></script>

			<title>{{.Title}}</title>
		</head>
		<body>
			<form hx-put="/contact/edit" hx-target="this" hx-swap="outerHTML">
				<div>
					<label>First Name</label>
					<input type="text" name="firstName" value="{{.FirstName}}">
				</div>
				<div class="form-group">
					<label>Last Name</label>
					<input type="text" name="lastName" value="{{.LastName}}">
				</div>
				<div class="form-group">
					<label>Email Address</label>
					<input type="email" name="email" value="{{.Email}}">
				</div>
				<button class="btn">Submit</button>
				<button class="btn" hx-get="/contact">Cancel</button>
			</form>
		</body>
	</html>` 
)
