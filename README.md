<div align="center">

# Twitter Chafa

**Go REST API that implements core Twitter/X features (users, tweets, followers, and timeline).**  
**API REST en Go que simula las funciones principales de Twitter/X (usuarios, tweets, seguidores y timeline).**

[ ![ Go ]( https://img.shields.io/badge/Go-1.20-00ADD8?style=flat-square&logo=go&logoColor=white ) ]( https://go.dev/ )
[ ![ MongoDB ]( https://img.shields.io/badge/MongoDB-4EA94B?style=flat-square&logo=mongodb&logoColor=white ) ]( https://www.mongodb.com/ )
[ ![ JWT ]( https://img.shields.io/badge/Auth-JWT-8A2BE2?style=flat-square&logo=jsonwebtokens&logoColor=white ) ]( https://jwt.io/ )
[ ![ License: MIT ]( https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square ) ]( https://opensource.org/licenses/MIT )
</div>

---

## What it does / Qué hace

- User registration and login with JWT authentication.  
  *Registro e inicio de sesión de usuarios con autenticación JWT.*

- Create, read, and delete tweets.  
  *Crear, leer y eliminar tweets.*

- Follow and unfollow users.  
  *Seguir y dejar de seguir usuarios.*

- View a user profile and update information (including avatar and banner).  
  *Ver el perfil de un usuario y editar datos (incluyendo avatar y banner).*

- Timeline showing tweets from the people the user follows.  
  *Timeline con los tweets de las personas que el usuario sigue.*

## Tech Stack / Tecnologías

- **Language / Lenguaje:** Go 1.20
- **Router:** Gorilla Mux
- **Database / Base de datos:** MongoDB
- **Auth / Autenticación:** JWT
- **CORS:** rs/cors

## How to run it / Cómo correrlo

### Requirements / Requisitos

- Go 1.20 or higher / Go 1.20 o superior
- MongoDB (local or Atlas) / MongoDB (local o Atlas)

### Steps / Pasos

```bash
git clone https://github.com/alanotonielgarciahernandez/Twitter_Chafa.git
cd Twitter_Chafa
go mod tidy
```

Configure the environment variables needed for the MongoDB connection (check `db/dbConnection.go` for expected values).  
*Configura las variables de entorno necesarias para la conexión a MongoDB (revisa `db/dbConnection.go` para ver los valores esperados).*

```bash
go run main.go
```

## Technical decisions / Decisiones técnicas

- MongoDB was chosen for its flexibility with user and tweet documents.  
  *Se usó MongoDB por su flexibilidad con documentos de usuarios y tweets.*

- Authentication is handled with JWT to protect endpoints that require a logged-in user.  
  *La autenticación se maneja con JWT para proteger los endpoints que requieren usuario logueado.*

- The project is organized in layers (routers, handlers, db, models) to keep the code clear.  
  *El proyecto está organizado por capas (routers, handlers, db, models) para mantener el código claro.*

## Status / Estado

Learning project, archived. No frontend or public demo.  
*Proyecto de aprendizaje archivado. No tiene frontend ni demo pública.*

## License / Licencia

MIT
