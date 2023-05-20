#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <time.h>

#define LICENSE_LENGTH 30
#define MAX_PATH 260

int check_license(char license[]);

int main() {
    DIR *d;
    struct dirent *dir;
    char license[LICENSE_LENGTH + 1] = "";
    const char version[10] = "7.5.0";
    const char path[MAX_PATH + 1] = "C:/Notepad";
    char expire[30] = "";
    int found = 0;

    if ((d = opendir("."))) {
        while ((dir = readdir(d)) != NULL) {
            if (strstr(dir->d_name, ".wad")) {
                FILE *file = fopen(dir->d_name, "r");
                if (file) {
                    found = 1;
                    fscanf(file, "Version=%s\nPath=%s\nExpire=%s", version, path, expire);
                    fclose(file);
                    strncpy(license, dir->d_name, strlen(dir->d_name) - 4);
                    license[strlen(dir->d_name) - 4] = '\0';
                    break;
                }
            }
        }
        closedir(d);
    }

    if (!found) {
        do {
            printf("Enter your WadBot license key: ");
            fgets(license, LICENSE_LENGTH + 2, stdin);
            if (license[strlen(license) - 1] == '\n') {
                license[strlen(license) - 1] = '\0';
            }
        } while (strlen(license) != LICENSE_LENGTH);

        time_t now = time(NULL);
        now += check_license(license);
        strftime(expire, sizeof(expire), "%Y-%m-%d %H:%M:%S", localtime(&now));

        char filename[LICENSE_LENGTH + 5];
        sprintf(filename, "%s.wad", license);
        FILE *file = fopen(filename, "w");
        if (file) {
            if (fprintf(file, "Version=%s\nPath=%s\nExpire=%s", version, path, expire) < 0) {
                printf("Error writing to file.\n");
            }
            fclose(file);
        } else {
            printf("Error opening file for writing.\n");
        }
    }

    return 0;
}

int check_license(char license[]) {
    return rand() % 10000;
}
