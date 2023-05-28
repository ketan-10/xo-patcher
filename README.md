# xo-patcher
Xo patcher is a tool used to generate sql script for bulk data migration.

## How it works 
- **XO** : Given the databse connection, we will read all your database schema, And create entites, and repositories for each table. 
- **Patcher**: Write your migration code in golang using entities and repositories created by XO, we will output `.sql` file containing all your mutation(update/insert) quries.

## How to use:
- `cd patcher && go mod tidy` Install dependencies.
- `cp .env.template .env` Then update enviroments according to you.
- `bash do.sh xo` This will generate entities and repositories by reading sql database given in `.env` file.
- `bash do.sh goimports` This will fix imports for generated code.
- `bash do.sh wire` This will perform dependency injection on the generated code.
- `touch patches/my-patch.go` In this file you can write your patch. 
- If you are using vscode type `init` to use the starup snippet.
- Register the patch in `wire_app/wire.go` so the patch can be injected in patchManager to be executed.
- `bash do.sh runlocal <patch_name>` to run the patch, generated `.sql` file will be placed in `\patches_gen`.
 