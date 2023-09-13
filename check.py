from datetime import datetime


def check_timestamps_in_order(filename):
    previous_timestamp = None

    with open(filename, "r") as file:
        for line in file:
            timestamp_str = line.strip()

            try:
                current_timestamp = datetime.strptime(timestamp_str, "%Y-%m-%d %H:%M:%S.%f")
            except ValueError:
                print(f"Invalid timestamp format in line: {timestamp_str}")
                return False

            if previous_timestamp is not None and current_timestamp < previous_timestamp:
                print(previous_timestamp, current_timestamp)
                return False

            previous_timestamp = current_timestamp

    return True


filename = "counter.txt"
result = check_timestamps_in_order(filename)

if result:
    print("Timestamps are in ascending order.")
else:
    print("Timestamps are not in ascending order.")
