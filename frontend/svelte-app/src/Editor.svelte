<script>
  import { onMount } from "svelte";
  import { writable } from "svelte/store";

  let fontSize = 16;
  let editor;
  let isVimEnabled = false;
  let isEmacsEnabled = false;
  let isAutoEnabled = false;
  let isOpen = false;
  let output;

  import "ace-builds/src-noconflict/ace";
  import "ace-builds/src-noconflict/theme-twilight";
  import "ace-builds/src-noconflict/keybinding-vim";
  import "ace-builds/src-noconflict/keybinding-emacs";
  import "ace-builds/src-noconflict/mode-text";
  import "ace-builds/src-noconflict/mode-golang";


  function toggleSettings() {
    isOpen = !isOpen;
  }

  const url = "/api/code";
  onMount(() => {

    editor = ace.edit("editor");

    editor.session.setMode("ace/mode/golang");

    editor.setTheme("ace/theme/twilight");

    editor.getSession().setValue("");

    editor.setBehavioursEnabled(true); 

  });

  function enableAutoBrackets() {
    isAutoEnabled = !isAutoEnabled;

    if (isAutoEnabled) {
      addCommand("autoClosingBrackets", editor, "[", "]")
      addCommand("autoClosingCurlyBraces", editor, "{", "}")
      addCommand("autoClosingParenthesis", editor, "(", ")")
    } else {

      editor.commands.removeCommand("autoClosingBrackets");
      editor.commands.removeCommand("autoClosingCurlyBraces");
      editor.commands.removeCommand("autoClosingParenthesis");
    }

  }

  function addCommand(name, editor, left, right) {
    editor.commands.addCommand({
      name: name,
      bindKey: { win: left, mac: left },
      exec: function (editor) {
        editor.insert(left + right);
        editor.navigateLeft();
      },
    });

  }

  function toggleEmacs() {
    isEmacsEnabled = !isEmacsEnabled;

    if (isEmacsEnabled) {
      isVimEnabled = false;
      editor.setKeyboardHandler("ace/keyboard/emacs");
    } else {
      editor.setKeyboardHandler(null);
    }
  }

  function toggleVim() {
    isVimEnabled = !isVimEnabled;

    if (isVimEnabled) {
      isEmacsEnabled = false;
      editor.setKeyboardHandler("ace/keyboard/vim");
    } else {
      editor.setKeyboardHandler(null);
    }
  }

  function readEditorContent() {
    return editor.getSession().getValue();
  }


  function sendCode() {
    const text = readEditorContent();

    const jsonObj = JSON.stringify(
        {
            text: text,
        }
    )
    fetch(url, {
      method: "POST",
      body: jsonObj,
    })
    .then(response => {
      if (response.ok) {
	    return response.json();
      }
    })
    .then(data => {
      const str = data.output.replace(/\n/g, '<br>');

      const err = data.err;
      console.log(err);
      if (err != undefined) {
        return;
      }
      output.innerHTML = str;
    });
}


  function increaseFontSize() {
    fontSize += 1;
  }

  function decreaseFontSize() {
    if (fontSize > 1) {
      fontSize -= 1;
    }
  }

</script>

<style>
  
  #editor {
    border-radius: 10px;
    height: 600px;
    width: 1250px;
    border-color: black;
    border-width: 1px;
  }
  .container {
    display: flex;
    width: 100%;
    flex-direction: column;
    border-radius: 10px;  
    align-items: center;
    justify-content: center;
  }
  .output {
    color: white;
    text-align:center;
    height: 100px;
    font-weight: bold;
    width: 300px;
    border-radius: 10px;
    border-width: 1px;
    border-color: black;
    background-color: rgb(27 27 27);
  }

  .settings {
    width: 20%;
    height: auto;
    display: none;
    border-width: 3px;
    border-style: solid;
    border-color: black;
    visibility: hidden;
    padding: 10px;

    flex-direction: column;
    justify-content: center;
    align-content: center;
  }
  .open {
    display: flex;
    visibility: visible;
  }

  .settings button{
    height: 50px;
    width: 100%;
    border-width: 2px;
    border-color: black;
  }

  .settings input{
    height: 50px;
    width: 100%;
    border-width: 2px;
    border-color: black;
  }

  .settings button:hover{
    border-color: blue;
  }

  .active {
    background-color: lightgreen;
    font-weight: bold;
  }

  .stBtn {
    border-radius: 10px;
    font-size: 30px;
  }


</style>

<button on:click={toggleSettings} class="stBtn">Settings: ⚙</button>

<div class="container">


  <div class="settings" class:open={isOpen}>
    <br>
    <label for="">Keybind Modes</label>
    <button on:click={toggleVim} class:active={isVimEnabled}> Toggle Vim </button>
    <button on:click={toggleEmacs} class:active={isEmacsEnabled}> Toggle Emacs </button>
    <br>
    <label for="">Settings</label>
    <button on:click={enableAutoBrackets} class:active={isAutoEnabled}> Toggle Auto Closing </button>
    <br>
    <label for="">Font Settings</label>
    <input bind:value={fontSize}/>
    <button on:click={increaseFontSize}>Increase</button>
    <button on:click={decreaseFontSize}>Decrease</button>
    <br>
  </div>

  <br>
  <br>
  <br>

    <div id="editor" style="font-size: {fontSize}px;" ></div>
    <br>
    <button on:click={sendCode}> Run Code </button>
    <br>
    <div class="output">
      <div bind:this={output}></div>
    </div>
</div>
