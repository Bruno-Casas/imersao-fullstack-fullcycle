#!/bin/bash

npm install
npm run typeorm migration:run
npm run console fixture
npm run start:dev