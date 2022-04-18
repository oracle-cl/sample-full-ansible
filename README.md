# Ejemplo Ansible
Despliega app web (golang)
## Parámetros
- globales.yaml: Generado por terraform con datos TNS
- lista_hosts: Generado por Terraform con las IP de las máquinas

## Ejecución
- Desde la máquina Bastión

## Playbook
- inicia.yaml: Configura TNS, Inicializa database schema, despliega y ejecuta aplicativo golang
