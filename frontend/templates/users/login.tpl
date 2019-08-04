{{ define "users_login" }}
<section>
  <article>
    <header class="article-header">
      <div class="row">
        <h1 class="display">Logga in</h1>
      </div>
    </header>
    <div class="row">
			<form style="width: 500px; margin: auto" method="post">
        {{ with .D.Error }}
          <div class="alert alert-error">{{ . }}</div>
        {{ end }}

				<div class="form-group">
					<div class="input-field">
						<label>Användarnamn</label>
						<input type="text" required name="username" placeholder="Fyll i ditt användarnamn">
					</div>
				</div>

				<div class="form-group">
					<div class="input-field">
						<label>Lösenord</label>
						<input type="password" required name="password" placeholder="Fyll i ditt lösenord">
					</div>
				</div>

				<div class="form-group">
					<div class="submit-field">
						<input type="submit" value="Logga in" class="raised">
					</div>
				</div>
			</form>
    </div>
  </article>
</section>
{{ end }}
