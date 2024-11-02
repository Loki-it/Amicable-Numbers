import math
import multiprocessing
import time

def sum_of_divisors(n):
    divisors_sum = 1
    sqrt_n = int(math.sqrt(n))
    for i in range(2, sqrt_n + 1):
        if n % i == 0:
            divisors_sum += i
            if i != n // i:
                divisors_sum += n // i
    return divisors_sum

def find_amicable_numbers_in_range(start, end):
    amicable_pairs = []
    for num in range(start, end):
        sum_div = sum_of_divisors(num)
        if sum_div > num and sum_of_divisors(sum_div) == num:
            amicable_pairs.append((num, sum_div))
    return amicable_pairs

def parallel_find_amicable_numbers(limit):
    num_processes = multiprocessing.cpu_count()  # Usa il numero massimo di core disponibili
    pool = multiprocessing.Pool(num_processes)
    chunk_size = limit // num_processes
    ranges = [(i * chunk_size, (i + 1) * chunk_size) for i in range(num_processes)]
    results = pool.starmap(find_amicable_numbers_in_range, ranges)
    pool.close()
    pool.join()
    amicable_pairs = [pair for sublist in results for pair in sublist]
    return amicable_pairs

if __name__ == '__main__':
    limit = 10**6  # Imposta il numero che vuoi testare
    start_time = time.time()
    amicable_pairs = parallel_find_amicable_numbers(limit)
    end_time = time.time()
    print(f"Coppie di numeri amici trovate fino a {limit}:")
    for pair in amicable_pairs:
        print(pair)

    print(f"\nTempo impiegato: {end_time - start_time} secondi")
