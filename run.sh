#!/bin/bash

set -e

vim -c 'so config.vim' -- $(cat list)
