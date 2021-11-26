import random
import sys

CHARS = "abcdefghijklmnopqrstuvwxyz"
NUMBERS = "0123456789"
ALPHANUM = CHARS + NUMBERS


def rand_word(max_len=5):
    word = random.choice(CHARS)
    for _ in range(random.randint(1, max_len)):
        word += random.choice(ALPHANUM)
    return word


def rand_protocol():
    return random.choice(["http", "https"])


def rand_port(none_chance=0.8):
    if random.random() > none_chance:
        return random.choice(["80", "8000", "8080", "5000", "443", "4200"])
    return ""


def random_times(f, min_times=0, max_times=5, none_chance=0.0) -> list:
    tokens = []
    if random.random() > none_chance:
        for i in range(random.randint(min_times, max_times)):
            tokens.append(f())
    return tokens


def random_url():
    port = rand_port()
    host_tokens = random_times(rand_word, 2, 4)
    host_tokens.append(random.choice(["com", "net", "ch", "jp", "ru", "us", "uk"]))
    path_tokens = random_times(rand_word, 1, 4, 0.4)
    parameters_tokens = random_times(rand_word, 0, 2)
    url = rand_protocol() + "://" + ".".join(host_tokens)
    if port:
        url += ":" + port
    if path_tokens:
        url += "/" + "/".join(path_tokens)
    if parameters_tokens:
        url += "?" if random.random() > 0.3 else "/?"
        url += "&".join(["{}={}".format(param, random.randint(0, 10000)) for param in parameters_tokens])
    return url


if __name__ == '__main__':
    if len(sys.argv) > 1:
        try:
            max_size = int(sys.argv[1])
        except Exception:
            max_size = 100
    else:
        max_size = 100

    for i in range(max_size):
        print(random_url())
