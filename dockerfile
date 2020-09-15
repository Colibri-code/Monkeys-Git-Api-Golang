#Base Image
FROM fedora:latest 

# Main commands that install packages dnf is fedoras package manager
RUN dnf install git \
    nodejs \
    python3 -y


#Specifices the port that is shared with ur main os
ENV PORT = 3000

#Shares the Port
EXPOSE 3000

# Specifies where the data that you write in the container persists after you exit the container.
# Data of the application goes in here.
VOLUME ["/DATA"]

