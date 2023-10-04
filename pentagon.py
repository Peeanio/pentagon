import pygame

BLACK = (0, 0, 0)
WHITE = (200, 200, 200)
WINDOW_HEIGHT = 500
WINDOW_WIDTH = 500
circle_pos = [40, 40]
circle2_pos = [-40, -40]
circle_selected = True

def main():
    running = True
    global SCREEN, CLOCK, circle_pos, circle_selected
    pygame.init()
    SCREEN = pygame.display.set_mode((WINDOW_WIDTH, WINDOW_HEIGHT))
    CLOCK = pygame.time.Clock()
    SCREEN.fill(BLACK)

    while running:
        drawGrid()
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            if event.type == pygame.MOUSEBUTTONDOWN:
                pos=pygame.mouse.get_pos()
                btn=pygame.mouse
                x = round(pos[0], -1)
                str_x = str(x)
                y = round(pos[1], -1)
                str_y = str(y)
                print ("x = {}, y = {}".format(round(x, -1), round(y, -1)))
                if int(str_x[:-1]) % 2 != 0:
                    print("round x")
                    x += 10
                if int(str_y[:-1]) % 2 != 0:
                    print("round y")
                    y += 10
                print ("x = {}, y = {}".format(round(x, -1), round(y, -1)))
                if x == circle_pos[0] and y == circle_pos[1] and circle_selected == False:
                    print("picked")
                    circle_selected = True
                    break
                if circle_selected == True:
                    print("placed")
                    circle_pos = [x, y]
                    circle_selected = False

        pygame.display.update()


def drawGrid():
    blockSize = 20 #Set the size of the grid block
    SCREEN.fill("black")

    for x in range(0, WINDOW_WIDTH, blockSize):
        for y in range(0, WINDOW_HEIGHT, blockSize):
            rect = pygame.Rect(x, y, blockSize, blockSize)
            pygame.draw.rect(SCREEN, WHITE, rect, 1)

    # pygame.draw.circle(SCREEN, "red", circle_pos, 10)
    natoImg = pygame.image.load('nato_small.png')
    SCREEN.blit(natoImg, circle_pos)
main()
