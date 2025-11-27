cmake_minimum_required(VERSION 3.10)
project(Lab3_OOP)

set(CMAKE_CXX_STANDARD 17)

# Флаги для покрытия кода (Coverage)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O0 -fprofile-arcs -ftest-coverage --coverage")
set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} --coverage")

# Подключаем папку tests, чтобы видеть catch.hpp
include_directories(${CMAKE_SOURCE_DIR})
include_directories(${CMAKE_SOURCE_DIR}/tests)

# --- 1. Google Test ---
include(FetchContent)
FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/refs/heads/main.zip
  DOWNLOAD_EXTRACT_TIMESTAMP TRUE
)
FetchContent_MakeAvailable(googletest)

# Сборка тестов Google Test (основные тесты для покрытия)
add_executable(test_runner_gtest tests/tests_gtest.cpp)
target_link_libraries(test_runner_gtest GTest::gtest_main)

# --- 2. Boost Test ---
# Ищем Boost (без указания компонентов, чтобы не было ошибки)
find_package(Boost REQUIRED)
include_directories(${Boost_INCLUDE_DIRS})

add_executable(test_runner_boost tests/tests_boost.cpp)

# --- 3. Catch2 ---
add_executable(test_runner_catch tests/tests_catch.cpp)

# --- Benchmark ---
add_executable(run_bench tests/bench_main.cpp)