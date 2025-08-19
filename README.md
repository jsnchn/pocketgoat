# PocketGoat

A Go application built with PocketBase and Templ for modern web development.

## Tech Stack

- **Backend**: Go 1.24+ with [PocketBase](https://pocketbase.io/) for rapid development
- **Templating**: [Templ](https://github.com/a-h/templ) for type-safe HTML templates
- **Frontend**: [Alpine.js](https://alpinejs.dev/) for reactive UI components
- **Styling**: [Tailwindcss](https://tailwindcss.com) and [DaisyUI](https://daisyui.com) for rapid ui development
- **Development**: [Air](https://github.com/air-verse/air) for hot reloading
- **Tool Management**: [mise](https://mise.jdx.dev/) for consistent environment tooling

## Prerequisites

### Option 1: Using mise (Recommended)

[mise](https://mise.jdx.dev/) is a polyglot tool version manager that ensures consistent development environments across team members.

```bash
# Install mise (if not already installed)
curl https://mise.run | sh

# Install all required tools automatically
mise install
```

### Option 2: Manual Installation

- Go 1.24 or higher
- [Air](https://github.com/air-verse/air) for hot reloading
- [Templ](https://github.com/a-h/templ) for templating

## Initial Setup

### 1. Install Dependencies

```bash
# Install Go dependencies
go mod download

# Install development tools
go install github.com/cosmtrek/air@latest
go install github.com/a-h/templ/cmd/templ@latest
```

### 2. Run the Application

```bash
# Development mode with hot reloading
air

# Or build and run manually
templ generate
go build -o ./tmp/main .
./tmp/main serve --http=127.0.0.1:8090
```

The application will be available at `http://127.0.0.1:8090`

### 3. Create PocketBase Superuser

On first run, PocketBase will prompt you to create a superuser account. This account will have full access to the PocketBase admin panel at `http://127.0.0.1:8090/_/`

### 4. Configure Google OAuth

After creating your superuser account:

1. Log in to the PocketBase admin panel
2. Navigate to **Settings** → **Auth providers**
3. Enable **Google** OAuth provider
4. Configure with your Google OAuth credentials:
   - **Client ID**: Your Google OAuth client ID
   - **Client Secret**: Your Google OAuth client secret

To obtain Google OAuth credentials:
1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Enable the Google+ API
4. Create OAuth 2.0 credentials
5. Add authorized redirect URI: `http://127.0.0.1:8090/api/oauth2-redirect`

### 5. Configure Transactional Email Provider

To enable email functionality (password resets, verification emails, etc.):

1. In the PocketBase admin panel, go to **Settings** → **Mail settings**
2. Configure your SMTP settings:
   - **SMTP server**: Your email provider's SMTP server
   - **Port**: Usually 587 for TLS or 465 for SSL
   - **Username**: Your SMTP username
   - **Password**: Your SMTP password
   - **Use TLS**: Enable for secure connection

Popular transactional email providers:
- [SendGrid](https://sendgrid.com/)
- [Mailgun](https://www.mailgun.com/)
- [Postmark](https://postmarkapp.com/)
- [Amazon SES](https://aws.amazon.com/ses/)

## Project Structure

```
.
├── main.go                 # Main application entry point
├── templates/              # Templ templates
│   ├── home.templ         # Template source files
│   └── home_templ.go      # Generated template code
├── pb_data/               # PocketBase data (gitignored)
├── pb_migrations/         # PocketBase migrations
├── pb_public/             # Static files served by PocketBase
└── tmp/                   # Temporary build files
```

## Development

The project uses Air for hot reloading during development. Any changes to `.go`, `.templ`, or other watched files will automatically rebuild and restart the application.

Template files (`.templ`) are automatically compiled to Go code before building.

### Frontend Architecture

The application uses a minimal custom CSS approach for styling, providing:
- Lightweight and fast loading times
- Full control over styling without framework overhead
- Responsive design with modern CSS features
- Clean, semantic HTML structure

Frontend interactivity is handled by Alpine.js, providing:
- Reactive data binding
- Component state management
- Seamless integration with server-rendered templates

## Production Deployment

For production deployment:

1. Build the application:
   ```bash
   templ generate
   go build -o pocketgoat .
   ```

2. Run with production settings:
   ```bash
   ./pocketgoat serve --http=0.0.0.0:8090
   ```

3. Use environment variables or flags to configure:
   - Database location
   - Server address
   - OAuth credentials
   - Email settings

Consider using a reverse proxy (nginx, Caddy) for SSL termination and additional security.

## License

MIT