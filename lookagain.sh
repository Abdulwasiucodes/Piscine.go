x#! /bin/bash
#!/bin/bash

# Find all .sh files, get only the filename, remove the .sh extension, and sort in descending order
find . -type f -name "*.sh" -exec basename {} \; | cut -d'.' -f1 | sort -r
