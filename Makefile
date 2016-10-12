CXX=g++
CXXFLAGS=-Wall

build: solution
	$(CXX) $(CXXFLAGS) solution.cpp -o solution

run: solution
	$(CXX) solution.cpp -o solution
	@echo "=========="
	@./solution
	@echo "\n=========="
	@rm -rf ./solution


clean:
	@rm -rf ./solution
