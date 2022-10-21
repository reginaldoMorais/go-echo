#!/bin/bash

cp ./tools/hooks/pre-commit ./.git/hooks
cp ./tools/hooks/pre-push ./.git/hooks

chmod +x ./.git/hooks/pre-commit
chmod +x ./.git/hooks/pre-push