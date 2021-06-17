const form = document.getElementById('form')
const input = document.getElementById('input')
const todosUL = document.getElementById('todos')

// baseURL
const baseURL = 'http://localhost/todos'
console.log(baseUrl);
// サーバーからデータを取得
let todos
getAllTodo()

// Todoを取得する
async function getAllTodo() {
    const response = await fetch(baseURL + 'todos')
    const todos = await response.json()
    // 画面を開いた時にリストを生成する
    if(todos) {
        todos.forEach(todo => addTodo(todo))
    }
}

// Todoを格納する
async function store(todo){
    await fetch(baseURL + 'todo/store',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(todo),
    })
}

// Todoのステータスを更新する
async function statusUpdate(id){
    await fetch(baseURL + 'todo/statusupdate',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({id: id}),
    })
}

// Todoを削除する
async function deleteTodo(id){
    await fetch(baseURL + 'todo/delete',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({id: id}),
    })
}

form.addEventListener('submit', (e) => {
    e.preventDefault()
    addTodo()
})

function addTodo(todo) {
    let todoText = input.value

    // 1 ~ 10000までのランダムな文字列
    let id = Math.floor( Math.random() * (10000 + 1 - 1) ) + 1

    // データを作成
    const todoData = {
        id: id,
        text: todoText
    }

    if(todo) {
        // idを入力
        id = todo.id
        todoText = todo.text
    }

    // 新規かつテキストに入力があるのみ
    if (!todo && todoText) {
        // バックエンドにアップロード
        store(todoData)
    }

    if(todoText) {

        const todoEl = document.createElement('li')
        if(todo && todo.completed) {
            todoEl.classList.add('completed')
        }

        todoEl.innerText = todoText

        todoEl.addEventListener('click', () => {
            todoEl.classList.toggle('completed')
            // ステータスアップデート
            statusUpdate(id)
        })

        todoEl.addEventListener('contextmenu', (e) => {
            e.preventDefault()
            todoEl.remove()
            // タスクデリート
            deleteTodo(id)
        })
        todosUL.appendChild(todoEl)
        input.value = ''
    }
}