cmake_minimum_required(VERSION 3.10)
project(Lab3_OOP)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Пакеты
find_package(GTest REQUIRED)
find_package(Boost REQUIRED)
include_directories(${Boost_INCLUDE_DIRS})

# Пути к заголовкам
include_directories(${CMAKE_CURRENT_SOURCE_DIR})
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/tests)

# Глобальные настройки покрытия
# Добавляем флаги компиляции и линковки
add_compile_options(-g -O0 --coverage)
add_link_options(--coverage)

# Цели

# 1. Google Test
add_executable(run_google tests/tests_google.cpp)
target_link_libraries(run_google GTest::gtest GTest::gtest_main pthread gcov)

# 2. Boost Test
add_executable(run_boost tests/tests_boost.cpp)
target_link_libraries(run_boost gcov)

# 3. Catch2
add_executable(run_catch tests/tests_catch.cpp)
target_link_libraries(run_catch gcov)

# 4. Benchmark
add_executable(run_bench tests/bench_main.cpp)
target_link_libraries(run_bench gcov)