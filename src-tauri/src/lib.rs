use tauri::menu::MenuBuilder;
use tauri::menu::SubmenuBuilder;

// Learn more about Tauri commands at https://tauri.app/develop/calling-rust/
#[tauri::command]
fn greet(name: &str) -> String {
    format!("Hello, {}! You've been greeted from Rust!", name)
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .setup(|app| {
            let submenu = SubmenuBuilder::new(app, "App").quit().build()?;

            let menu = MenuBuilder::new(app).items(&[&submenu]).build()?;

            app.set_menu(menu)?;

            Ok(())
        })
        .plugin(tauri_plugin_sql::Builder::new().build())
        .plugin(tauri_plugin_pinia::init())
        .plugin(tauri_plugin_opener::init())
        .invoke_handler(tauri::generate_handler![greet])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
