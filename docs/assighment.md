1. Repozytorium

   stwórz nowe repozytorium na GitHubie
   nazwij je: idp-jobs-go
   ustaw jako publiczne
   dodaj mnie jako collaborator: https://github.com/norbix/

1. Zadanie techniczne

    Zaimplementuj endpoint:
    
    GET /idp/jobs/{jobId}
    
    który zwraca:
    
        jobId
        status
        createdAt

1. Wymagania

   podział na warstwy:
   - api
   - service
   - repository
   - model
   - storage:
   - na start in-memory
   
   testy:
   - unit test dla service
   - integration test dla endpointu HTTP z uzyciem docker

1. Dodatkowo

   README.md z:
   jak uruchomić projekt
   jak uruchomić testy

