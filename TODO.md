TODO list

- Tool to delete log files en la web (tipo ftp para verlos y borrarlos?)
- Lista negra de hostnames
- Ampliar lista de STATUS
- Encriptar contraseñas
- Propuesta: On management server startup, crear usuario admin si no existe, con una contraseña random. 
Este usuario y contraseña será el de login de la web de management. 
Se debe comprobar la constraseña directamente en la bbdd mongo (también protegida con usuario y contraseña). 
Se recomienda cambiar la contraseña después de iniciar sesion por primera vez.
Este usuario admin no puede borrarse