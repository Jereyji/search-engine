import requests
from bs4 import BeautifulSoup

# Функция для парсинга сайта NGS.ru
def parse_ngs():
    url = 'https://ngs.ru/'
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')

    # Ищем блоки с заголовками новостей
    headlines = soup.find_all('a', class_='news-title')

    news_list = []
    for item in headlines:
        title = item.get_text(strip=True)
        link = item['href']
        if not link.startswith('http'):
            link = 'https://ngs.ru' + link
        news_list.append((title, link))

    return news_list

# Функция для парсинга сайта Gazeta.ru
def parse_gazeta():
    url = 'https://gazeta.ru/'
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')

    # Ищем блоки с заголовками новостей
    headlines = soup.find_all('a', class_='news-item__title')

    news_list = []
    for item in headlines:
        title = item.get_text(strip=True)
        link = item['href']
        if not link.startswith('http'):
            link = 'https://gazeta.ru' + link
        news_list.append((title, link))

    return news_list

# Функция для парсинга сайта Lenta.ru
def parse_lenta():
    url = 'https://lenta.ru/'
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')

    # Ищем блоки с заголовками новостей
    headlines = soup.find_all('a', class_='card-mini__title')

    news_list = []
    for item in headlines:
        title = item.get_text(strip=True)
        link = item['href']
        if not link.startswith('http'):
            link = 'https://lenta.ru' + link
        news_list.append((title, link))

    return news_list

# Основная функция, которая вызывает парсеры и выводит новости
def main():
    print("Новости с NGS.ru:")
    for title, link in parse_ngs():
        print(f"{title}: {link}")

    print("\nНовости с Gazeta.ru:")
    for title, link in parse_gazeta():
        print(f"{title}: {link}")

    print("\nНовости с Lenta.ru:")
    for title, link in parse_lenta():
        print(f"{title}: {link}")

if __name__ == "__main__":
    main()
