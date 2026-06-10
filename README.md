# 机场贵宾厅预约权益核验端到端协同控制台

## 项目简介

这是一个面向机场贵宾厅预约权益核验业务的中等复杂度全栈应用系统，专为**权益运营人员、机场服务台工作人员和客服支持团队**设计。

### 核心功能

- ✅ 会员权益创建与流转管理
- ✅ 核验权益资格前端工作台
- ✅ 候补入场后端计算与校验
- ✅ 权益过期异常分诊处理
- ✅ 航班时段详情与历史记录
- ✅ 权益核验通过率统计驾驶舱
- ✅ 同行人批量导入和复核
- ✅ 规则配置、命中解释和审计日志
- ✅ 归档快照与数据导出
- ✅ Seed数据浏览和接口健康检查

## 技术栈

### 前端 (Frontend)
- **框架**: Svelte 4.x + TypeScript
- **构建工具**: Vite 5.x
- **状态管理**: Svelte Stores
- **表格组件**: TanStack Table v8
- **图表库**: Chart.js 4.x + svelte-chartjs
- **UI组件**: 自研高级组件（8个）
- **测试框架**: Vitest + jsdom

### 后端 (Backend)
- **语言**: Go 1.21+
- **Web框架**: Fiber v2
- **数据库**: SQLite3
- **架构模式**: Repository/DAO + Service + Handler 分层
- **测试框架**: Go标准testing包

## 快速开始

### 环境要求

- **Go**: 1.21 或更高版本
- **Node.js**: 18.x 或更高版本
- **npm**: 9.x 或更高版本

### 1. 克隆项目

```bash
cd /path/to/repo_0143
```

### 2. 安装并启动后端 (Backend)

#### 安装依赖

```bash
cd backend
go mod download
```

#### 初始化数据库

```bash
# 初始化SQLite数据库并运行迁移脚本
go run src/main.go --init-db
```

#### 导入Seed数据

```bash
# 导入50+条真实业务场景的种子数据
go run src/main.go --seed
```

#### 启动后端服务

```bash
# 启动后端服务器（默认端口8080）
go run src/main.go
```

后端服务启动成功后会显示：
```
✈️ Airport VIP Lounge Console - Backend Server
🚀 Server is running on http://localhost:8080
📊 API Base URL: http://localhost:8080/api/v1
🌱 Seed data loaded: XX records
```

### 3. 安装并启动前端 (Frontend)

#### 安装依赖

```bash
cd frontend
npm install
```

#### 启动开发服务器

```bash
# 启动前端开发服务器（默认端口5173）
npm run dev
```

前端开发服务器会自动打开浏览器访问 `http://localhost:5173`，并代理API请求到后端 `http://localhost:8080`。

### 4. 访问系统

在浏览器中打开：
- **前端界面**: http://localhost:5173
- **后端API文档**: http://localhost:8080/api/v1/health

## 核心功能使用指南

### 1. 统计驾驶舱 (Dashboard)

访问路径：侧边栏 → 📊 统计驾驶舱

展示内容：
- 总核验次数、通过率、候补人数等关键指标
- 核验结果分布饼图（通过/未通过/候补中）
- 权益使用情况分析
- 近7日趋势折线图

### 2. 会员权益管理

访问路径：侧边栏 → 💳 会员权益

支持操作：
- **新建权益**: 点击"➕ 新建权益"按钮填写表单
- **编辑修改**: 点击"✏️"图标编辑现有记录
- **状态流转**: 通过/驳回/归档等状态变更
- **批量操作**: 多选后进行批量归档、激活或删除
- **高级筛选**: 按状态、等级、机场等多维度筛选
- **命令面板**: 按 `Ctrl+K` 快速执行常用操作

### 3. 键盘快捷键

| 快捷键 | 功能 |
|--------|------|
| `Ctrl+K` | 打开命令面板 |
| `Ctrl+N` | 新建会员权益 |
| `F5` | 刷新当前列表 |

## API 接口说明

### RESTful API 列表 (12+个端点)

#### 健康检查 & Seed
- `GET /api/v1/health` - 健康检查
- `POST /api/v1/seed/reset` - 重置Seed数据

#### 会员权益 CRUD
- `GET /api/v1/member-benefits` - 获取会员权益列表（支持分页）
- `POST /api/v1/member-benefits` - 创建会员权益
- `PUT /api/v1/member-benefits/:id` - 更新会员权益
- `DELETE /api/v1/member-benefits/:id` - 删除会员权益
- `POST /api/v1/member-benefits/:id/transition` - 状态流转

#### 预约记录
- `GET /api/v1/appointments` - 获取预约记录
- `POST /api/v1/appointments/:id/transition` - 预约状态流转

#### 航班时段
- `GET /api/v1/flight-slots` - 获取航班时段
- `GET /api/v1/flight-slots/airport/:code` - 按机场查询

#### 同行人
- `GET /api/v1/companions` - 获取同行人列表
- `POST /api/v1/companions/batch` - 批量导入同行人

#### 候补名单
- `GET /api/v1/waiting-list` - 获取候补名单
- `POST /api/v1/waiting-list/process` - 处理候补入场

#### 统计聚合
- `GET /api/v1/statistics` - 获取统计数据（支持日期范围）

#### 领域计算
- `GET /api/v1/domain-calculations/priority/:id` - 计算优先级
- `GET /api/v1/domain-calculations/consistency/:id` - 一致性校验
- `GET /api/v1/domain-calculations/expired-benefits` - 检测过期权益

#### 异常事件
- `GET /api/v1/exceptions` - 获取异常事件列表

#### 审计日志
- `GET /api/v1/audit-logs` - 获取审计日志流

### 请求样本

**获取会员权益列表**
```http
GET /api/v1/member-benefits?page=1&page_size=10 HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

**响应样本**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "code": "VIP-BENEFIT-20260101-001",
      "name": "2026年度白金卡贵宾厅权益",
      "status": "active",
      "member_level": "platinum",
      "remaining_quota": 12,
      "valid_from": "2026-01-01T00:00:00Z",
      "valid_until": "2026-12-31T23:59:59Z",
      "owner": "张明",
      "batch_id": "BATCH-202601",
      "airport_code": "PEK",
      "lounge_id": "PEK-T3-DOMESTIC-VIP",
      "notes": "",
      "created_at": "2026-01-15T08:30:00Z",
      "updated_at": "2026-01-20T14:22:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "page_size": 10,
    "total_count": 50,
    "total_pages": 5
  }
}
```

## 运行测试

### 后端测试

```bash
cd backend
go test ./tests/... -v
```

后端测试包括：
- 数据库连接测试
- Seed数据加载测试
- API端点集成测试
- 业务规则验证测试

### 前端测试

```bash
cd frontend

# 运行所有测试
npm test

# 运行测试并生成覆盖率报告
npm run test:coverage

# 使用UI界面查看测试结果
npm run test:ui
```

前端测试包括：
- API Service方法定义测试
- 前后端集成测试（需后端运行中）
- 组件渲染测试（可扩展）

### 接口契约测试

集成测试会验证：
- ✅ Seed数据从后端正确返回
- ✅ 前端API调用与后端路由一致
- ✅ 字段映射和类型匹配
- ✅ 业务数据格式符合预期

## 核心业务规则

### 规则1：入场优先级计算
根据以下因素综合计算：
- 会员等级权重（白金>金卡>银卡>标准）
- 权益剩余额度（越多越优先）
- 航班时间紧迫度
- 候补等待时长

### 规则2：权益过期异常处理
- 自动检测过期权益
- 进入异常队列
- 记录触发字段、阈值、实际值
- 设置处理人和时限

### 规则3：一致性校验
- 核验前必须验证：会员权益 ↔ 预约记录 ↔ 航班时段
- 缺少关键字段时只允许保存为草稿
- 完整信息才能提交复核

### 规则4：统计指标聚合
支持按以下维度聚合：
- 日/周/月维度
- 批次号维度
- 责任角色维度
- 可钻取到明细数据源

### 规则5：状态机控制
允许的状态流转：
```
草稿 → 待复核 → 已确认 → 已归档
                  ↘ 已驳回（必须填写原因）

待补充 → 待复核 → 已确认
已确认 → 已归档
活跃 → 已过期（系统自动）
```

## Seed 数据说明

系统包含 **50+ 条真实业务场景** 的种子数据：

### 数据覆盖
- ✅ 正常状态记录（active）
- ✅ 异常事件记录（expired, rejected）
- ✅ 待处理记录（pending_review, draft）
- ✅ 已归档记录（archived）
- ✅ 不同会员等级分布
- ✅ 多个机场代码（PEK, PVG, CAN, SZX, CTU）

### 编号规范
- 会员权益：`VIP-BENEFIT-{日期}-{序号}`
- 预约记录：`APT-{日期}-{航班号}-{序号}`
- 航班时段：`SLOT-{机场}-{登机口}-{时间段}`
- 同行人：`COMP-{预约ID}-{序号}`
- 候补名单：`WAIT-{日期}-{序号}`

### 样本数据
```
VIP-BENEFIT-20260101-001: 张明的白金卡权益（活跃）
VIP-BENEFIT-20250615-042: 李娜的金卡权益（已过期）
APT-20260120-CA1234-003: CA1234航班的预约（待复核）
SLOT-PEK-GATE25-MORNING: 北京T3航站楼25号门早间时段
```

## 常见问题排查

### 问题1：后端启动失败 - 数据库锁定

**症状**:
```
database is locked
```

**解决方案**:
```bash
# 删除旧数据库文件

# 重新初始化
go run src/main.go --init-db --seed
```

### 问题2：前端无法连接后端API

**症状**:
```
Network Error / CORS error
```

**排查步骤**:
1. 确认后端正在运行：访问 http://localhost:8080/api/v1/health
2. 检查vite.config.ts中的proxy配置是否正确
3. 确认端口未被占用（后端8080，前端5173）

### 问题3：npm install 失败

**症状**:
```
npm ERR! code ERESOLVE
```

**解决方案**:
```bash
# 清除缓存
npm cache clean --force

# 删除node_modules和lock文件
rm -rf node_modules package-lock.json

# 重新安装
npm install
```

### 问题4：Go依赖下载失败

**症状**:
```
module xxx: Get "https://proxy.golang.org/...": dial tcp...
```

**解决方案**:
```bash
# 设置Go模块代理（国内用户）
export GOPROXY=https://goproxy.cn,direct

# 重新下载依赖
go mod download
```

### 问题5：测试失败 - jsdom缺失

**症状**:
```
Cannot find module 'jsdom'
```

**解决方案**:
```bash
cd frontend
npm install jsdom --save-dev
```

### 问题6：端口被占用

**症状**:
```
bind: address already in use
```

**解决方案**:
```bash
# 查找占用端口的进程
netstat -ano | findstr :8080  # Windows
lsof -i :8080                 # macOS/Linux

# 结束进程或更改端口
# 在backend/src/main.go中修改端口号
# 在frontend/vite.config.ts中修改proxy target
```

## 构建生产版本

### 前端构建

```bash
cd frontend
npm run build
```

构建产物位于 `frontend/dist/` 目录。

### 后端编译

```bash
cd backend
go build -o lounge-console.exe src/main.go
```

## 开发建议

### 推荐的开发工具
- **IDE**: VS Code + Svelte for VS Code扩展
- **浏览器**: Chrome/Firefox DevTools
- **API调试**: Postman / Insomnia
- **数据库查看**: DB Browser for SQLite

### 代码风格
- 前端遵循 ESLint + Prettier 配置
- 后端遵循 Go 官方代码规范
- 提交前运行 `npm run lint` 和 `go vet ./...`

## 项目特色

✨ **真实业务场景**: 不是简单的CRUD模板，而是围绕机场贵宾厅业务的完整工作流
✨ **领域驱动设计**: 包含优先级算法、状态机、异常检测等业务规则
✨ **前后端联调**: 所有API接口经过契约测试验证
✨ **丰富的UI交互**: 8个高级组件，支持键盘快捷键、拖拽、图表钻取
✨ **完善的测试覆盖**: 单元测试 + 集成测试 + 契约测试
✨ **可追溯的数据**: 统计结果可钻取到明细记录

## Notice

详见 [NOTICE](NOTICE) 文件

---

**开发完成时间**: 2026年05月25日
**技术支持**: 请查阅本文档或查看代码注释
