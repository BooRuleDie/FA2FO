.PHONY: clean
clean:
	@echo "Cleaning cache and temporary files..."
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type d -name "*.egg-info" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete
	rm -rf .pytest_cache
	rm -rf build/
	rm -rf dist/
	rm -rf .venv/lib/*/site-packages/__pycache__/
	rm -rf .venv/bin/__pycache__/
	rm -rf .venv
	pip cache purge
	@echo "Clean complete."

.PHONY: setup
setup:
	@echo "Installing Facebook Business SDK..."
	pip install facebook-business
	@echo "Installation complete."

.PHONY: freeze
freeze:
	@echo "Freezing requirements..."
	pip freeze > requirements.txt
	@echo "Requirements updated in requirements.txt"

.PHONY: restart
restart: clean setup freeze
