all process for develop
======================================

## create frontend project 
```commandline
npx create-react-app pizza-web
cd pizza-web
yarn start
```

## run mock API server
all the `json` data defined in file `db.json`

```commandline
npx json-server --watch db.json --port 5000
```

define the proxy in `pack`, visit the mock API code like

```javascript
fetch("/pizzas").then(response => response.json())
    .then(data => console.log(data));
```

## run python backend
```commandline
cd backend
python main.py
```
backend API will run in potr 5000







