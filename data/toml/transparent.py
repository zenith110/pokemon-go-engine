from PIL import Image, Image
import os 
import glob
def transparent_sprite(folder_name: str) -> None:
    # Loops through the entire folder
    folder_directory = os.listdir(folder_name)
    for files in folder_directory:
        image = Image.open(f"{folder_name}/{files}")
        rgba = image.convert("RGBA")
        datas = rgba.getdata()

        most_frequent_color = max(image.getcolors(image.size[0]*image.size[1]))
        try:
            newData = [
            most_frequent_color[1][0],
            most_frequent_color[1][1],
            most_frequent_color[1][2]
            ]
            
            imgData = []

            for item in datas:
                if item[0] == newData[0] and item[1] == newData[1] and item[2] == newData[2]:
                    imgData.append((newData[0], newData[1], newData[2], 0))
                else:
                    imgData.append(item)

            rgba.putdata(imgData)
            rgba.save(f"{folder_name}/{files}", "png")
        except Exception as e:
            print(f"Skipping {files}")
            print(e)
transparent_sprite("../assets/pokemon/front")
transparent_sprite("../assets/pokemon/back")
transparent_sprite("../assets/pokemon/shinyfront")
transparent_sprite("../assets/pokemon/shinyback")
transparent_sprite("../assets/pokemon/icons")