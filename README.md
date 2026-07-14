# Map data to Graphviz DOT language

## Usage
```
cd [this repo]
chmod +x map-to-dot.x86_64 
./map-to-dot.x86_64 [input map file] [output file]
```
You can use [Graphviz Online](https://dreampuf.github.io/GraphvizOnline/) to visualize the output.  
>[!NOTE]
>
>Input is read as plain text but file extension does not matter.  

## Example

**Input:**
```
stations:
# south stations
waterloo  , 3 , 1
victoria,6,7

# north stations
euston,11,23
st_pancras,5,15 # international


connections:
waterloo- euston
st_pancras-euston
victoria-st_pancras
waterloo -victoria
```
**Output:**
```
graph stations
{rankdir=LR;

node [shape=ellipse, fixedsize=false];

"waterloo";
"victoria";
"euston";
"st_pancras";

"waterloo" -- "euston";
"st_pancras" -- "euston";
"victoria" -- "st_pancras";
"waterloo" -- "victoria";

}
```
