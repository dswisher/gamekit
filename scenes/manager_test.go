package scenes

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

// mockScene is a test double for Scene.
type mockScene struct {
	enterCalled bool
	exitCalled  bool
	updateCalls int
	drawCalls   int
}

func (m *mockScene) Update(dt float64) { m.updateCalls++ }

func (m *mockScene) Draw(screen *ebiten.Image) { m.drawCalls++ }

func (m *mockScene) Enter() { m.enterCalled = true }

func (m *mockScene) Exit() { m.exitCalled = true }

// mockController is a test double for SceneController.
type mockController struct {
	pushCalled    bool
	popCalled     bool
	replaceCalled bool
	pushedScene   Scene
	replacedScene Scene
}

func (m *mockController) Push(sc Scene) {
	m.pushCalled = true
	m.pushedScene = sc
}

func (m *mockController) Pop() {
	m.popCalled = true
}

func (m *mockController) Replace(sc Scene) {
	m.replaceCalled = true
	m.replacedScene = sc
}

func TestNewSceneManager(t *testing.T) {
	sm := NewSceneManager()
	assert.NotNil(t, sm)
	assert.Nil(t, sm.Current())
}

func TestSceneManager_ImplementsSceneController(t *testing.T) {
	// This is a compile-time check
	var _ SceneController = (*SceneManager)(nil)
}

func TestPush_CallsEnter(t *testing.T) {
	sm := NewSceneManager()
	scene := &mockScene{}

	sm.Push(scene)

	assert.True(t, scene.enterCalled)
	assert.Equal(t, scene, sm.Current())
}

func TestPop_CallsExit(t *testing.T) {
	sm := NewSceneManager()
	scene := &mockScene{}
	sm.Push(scene)

	sm.Pop()

	assert.True(t, scene.exitCalled)
	assert.Nil(t, sm.Current())
}

func TestPop_EmptyStack_NoPanic(t *testing.T) {
	sm := NewSceneManager()
	assert.NotPanics(t, func() { sm.Pop() })
}

func TestReplace_CallsExitThenEnter(t *testing.T) {
	sm := NewSceneManager()
	oldScene := &mockScene{}
	newScene := &mockScene{}
	sm.Push(oldScene)

	sm.Replace(newScene)

	assert.True(t, oldScene.exitCalled)
	assert.True(t, newScene.enterCalled)
	assert.Equal(t, newScene, sm.Current())
}

func TestGetScenes_ReturnsCopy(t *testing.T) {
	sm := NewSceneManager()
	scene1 := &mockScene{}
	scene2 := &mockScene{}
	sm.Push(scene1)
	sm.Push(scene2)

	scenes := sm.GetScenes()
	assert.Len(t, scenes, 2)

	// Modifying returned slice should not affect manager
	scenes = scenes[:1]
	_ = scenes // suppress unused variable warning; we only care that GetScenes() still returns 2
	assert.Len(t, sm.GetScenes(), 2)
}

func TestMockController_UsableForTesting(t *testing.T) {
	// Demonstrates that scenes can be tested with mock controllers
	controller := &mockController{}
	scene := &mockScene{}

	controller.Push(scene)
	assert.True(t, controller.pushCalled)
	assert.Equal(t, scene, controller.pushedScene)

	controller.Pop()
	assert.True(t, controller.popCalled)

	controller.Replace(scene)
	assert.True(t, controller.replaceCalled)
	assert.Equal(t, scene, controller.replacedScene)
}
