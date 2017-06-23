const express = require('express');
const bodyParser = require('body-parser');

const app = express();
app.use(bodyParser.json());

app.post('/', (request, response) => {
    response.send('Hello '+request.body.name);
});

app.listen(3000);