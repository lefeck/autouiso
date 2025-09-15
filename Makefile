.PHONY: build run test clean help

# 默认目标
.DEFAULT_GOAL := help

# 变量定义
BINARY_NAME := autouiso
BUILD_DIR := build
MAIN_FILE := main.go

# 构建应用
build:
	@echo "构建 AutoISO 应用..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "构建完成: $(BUILD_DIR)/$(BINARY_NAME)"

# 运行应用
run: build
	@echo "启动 AutoISO 服务..."
	@$(BUILD_DIR)/$(BINARY_NAME)

# 开发模式运行（热重载）
dev:
	@echo "开发模式运行 AutoISO..."
	@go run $(MAIN_FILE)

# 运行测试
test:
	@echo "运行测试..."
	@go test ./...

# 清理构建文件
clean:
	@echo "清理构建文件..."
	@rm -rf $(BUILD_DIR)
	@go clean

# 安装依赖
deps:
	@echo "安装依赖..."
	@go mod tidy
	@go mod download

# 代码格式化
fmt:
	@echo "格式化代码..."
	@go fmt ./...

# 代码检查
lint:
	@echo "检查代码..."
	@go vet ./...

# 生成API文档
docs:
	@echo "生成API文档..."
	@mkdir -p docs
	@echo "# AutoISO API 文档" > docs/api.md
	@echo "" >> docs/api.md
	@echo "## 接口列表" >> docs/api.md
	@echo "" >> docs/api.md
	@echo "### 健康检查" >> docs/api.md
	@echo "\`\`\`" >> docs/api.md
	@echo "GET /api/v1/health" >> docs/api.md
	@echo "\`\`\`" >> docs/api.md

# 创建发布版本
release: clean build
	@echo "创建发布版本..."
	@tar -czf $(BINARY_NAME)-$(shell date +%Y%m%d).tar.gz $(BUILD_DIR)/

# 显示帮助信息
help:
	@echo "AutoISO 构建工具"
	@echo ""
	@echo "可用命令:"
	@echo "  build    构建应用"
	@echo "  run      构建并运行应用"
	@echo "  dev      开发模式运行（热重载）"
	@echo "  test     运行测试"
	@echo "  clean    清理构建文件"
	@echo "  deps     安装依赖"
	@echo "  fmt      格式化代码"
	@echo "  lint     代码检查"
	@echo "  docs     生成API文档"
	@echo "  release  创建发布版本"
	@echo "  help     显示此帮助信息"
	@echo ""
	@echo "示例:"
	@echo "  make build    # 构建应用"
	@echo "  make run      # 构建并运行"
	@echo "  make dev      # 开发模式运行"

