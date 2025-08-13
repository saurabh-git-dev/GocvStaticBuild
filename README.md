# Info
Repository contains a minimal example of opencv static linking in golang app with gocv.

# Gocv Static Build

1. Install Packages
    ```shell
    go mod tidy
    ```
2. Install Opencv 
    ```shell
    cd C:\Users\{USER}\go\pkg\mod\gocv.io\x\gocv@v0.42.0
    win_build_opencv.cmd static
    ```
3. Run build script
    ```
    bash build.sh
    ```

NOTE: I have included most of the basic dependencies needed to build executable. You may need add more deps depending upon you setup.

# How to know which deps required...
I will share my method here...
I started building executable with static flags. It will throw errors. Try to find out the cpp files names and try to relate with the static files in dir C:\\opencv\\build\\install\\x64\\mingw\\staticlib. If you find a anything similar. Add them to the deps list. Now Run again you may see another errors. Try to repeat the steps.
You may also encounter deps order related issues. This can be slight tricky but if you take the help of ChatGPT it might solve your issues. It may not suggest correct things. But you will get an idea of what is missing or what is the problem.
Otherwise you can also try to analyze error manually and include the deps in the order needed.
