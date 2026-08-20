# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

后台提醒发送失败后没有进入计划中的重试，而是直接落入永久失败队列；日志里原本的网络超时也变成了普通文本。此次只定位原因，不要改动目标仓库代码。请追踪可重试错误在各层传递时的身份变化，说明分类为何走错，并以日志或运行证据验证。

## 含 Bug 版本

- 仓库：11DingKing/chargeguard-bug-23
- 仓库地址：https://github.com/11DingKing/chargeguard-bug-23.git
- parent SHA：e5ed3ca593fbf2fa60111111c976ae0e70e13713

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/chargeguard-bug-23.git bug-repro
cd bug-repro
git checkout --detach e5ed3ca593fbf2fa60111111c976ae0e70e13713
go test ./internal/httpapi -run TestTaskBehavior -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:14: status=422 body=permanent failure
FAIL
FAIL	chargeguard/internal/httpapi	0.062s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/httpapi -run TestTaskBehavior -count=1
--- FAIL: TestTaskBehavior (0.00s)
    task_behavior_test.go:14: status=422 body=permanent failure
FAIL
FAIL	chargeguard/internal/httpapi	0.003s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

根因结论必须准确写明出问题的 Go 文件、具体符号和完整失效机制，并由实际复现、源码调查和验证证据支撑；调查结束时目标仓库代码、测试和配置零改动。
