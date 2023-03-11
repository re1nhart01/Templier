# Templier
 Template creator for components

## Flags:
```
-l - flag to set labels which define in !{ }! brackets
-n - flag to set file names for files inside @ @ brackets
--log - flag to set logger on, which will log every step
--help - current flag
-f - filename of file with components	
-p - pathname, where we need to create files
-c 1:3 - generate files from 1 to 3
-c 2   - generate file with count 2
```


## Example of templier.yaml snippet:

```
---
component2254:
  files:
  - name: "@text@.component.tsx"
    content:
      labels:
        componentName: labelName
        componentPropType: labelName
      data: |
          import React from 'react';
          
          type !{componentPropType}!componentType = {};
          
          const !{componentName}! = () => {
            return (
              <div>
                SomeText;
              </div>
            )
          }
  - name: "@text@.styles.tsx"
    content:
      labels:
        componentName: labelName
        componentPropType: labelName
      data: |
          import React from 'react';
          import { Stylesheet } from 'react-native';

          export default Stylesheet.create({

          })
  - name: "@text@.controller.tsx"
    content:
      labels: 
        - componentPropType
        - componentName
      data: |
          import React from 'react';
          
          type !{componentPropType}!componentType = {};
          
          const !{componentName}! = () => {
            return (
              <div>
                SomeText;
              </div>
            )
          }       
```



## Info
##### Welcome to Templier.
	The purpose of this application to generate template files.
	For example you have angular components.
	And you dont want to always create files,
	you can just define config yaml file and call this
	application with component name.
    Which are a default routine in developers live.

##### Created by @re1nhart
##### Written in Go programming language