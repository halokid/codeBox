#!/bin/bash

NEW_PASSWORD=$1
LOGIN_USER=$2
echo $NEW_PASSWORD | sudo passwd --stdin $LOGIN_USER

