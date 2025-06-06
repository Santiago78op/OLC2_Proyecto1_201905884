class IDEController {
    constructor() {
        this.fileManager = null;
        this.editor = null;
        this.reportsManager = null;
        this.isConnected = false;
        this.init();
    }

    init() {
        // Inicializar componentes
        this.fileManager = new FileManager();
        this.editor = new CodeEditor();
        this.reportsManager = new ReportsManager();

        // Hacer disponible globalmente
        window.ideController = this;
        window.reportsManager = this.reportsManager; // Para las funciones de ordenamiento

        // Configurar interfaz
        this.bindEvents();
        this.checkConnection();

        // Verificar conexión periódicamente
        setInterval(() => this.checkConnection(), 5000);

        this.addConsoleMessage('VLan Cherry IDE iniciado correctamente', 'info');
    }

    bindEvents() {
        // Botón ejecutar
        document.getElementById('executeBtn').addEventListener('click', () => {
            this.executeCode();
        });

        // Limpiar consola
        document.getElementById('clearConsoleBtn').addEventListener('click', () => {
            this.clearConsole();
        });

        // Eventos del menú de Electron
        if (window.electronAPI) {
            window.electronAPI.onMenuExecute(() => this.executeCode());
            window.electronAPI.onMenuShowReports(() => this.reportsManager.showReportsModal());
            window.electronAPI.onMenuClearConsole(() => this.clearConsole());
        }

        // Evento antes de cerrar ventana
        window.addEventListener('beforeunload', (e) => {
            if (this.fileManager.hasUnsavedChanges()) {
                e.preventDefault();
                e.returnValue = '💾 Hay cambios sin guardar. ¿Estás seguro de que quieres salir?';
            }
        });
    }

    switchPanel(panelName) {
        // Actualizar tabs
        document.querySelectorAll('.panel-tab').forEach(tab => {
            tab.classList.toggle('active', tab.dataset.panel === panelName);
        });

        // Mostrar/ocultar paneles
        document.querySelectorAll('.panel-content').forEach(panel => {
            panel.classList.toggle('active', panel.id === panelName + 'Panel');
        });
    }

    async checkConnection() {
        try {
            const response = await window.apiClient.checkStatus();
            this.updateConnectionStatus('🟢 Conectado', true);
            this.isConnected = true;
        } catch (error) {
            this.updateConnectionStatus('🔴 Desconectado', false);
            this.isConnected = false;
        }
    }

    updateConnectionStatus(message, isConnected) {
        const statusElement = document.getElementById('connectionStatus');
        statusElement.textContent = message;
        statusElement.className = `connection-status ${isConnected ? 'connected' : 'disconnected'}`;
    }

    async executeCode() {
        const activeFile = this.fileManager.getActiveFile();
        if (!activeFile) {
            this.addConsoleMessage('No hay archivo activo para ejecutar', 'warning');
            return;
        }

        if (!this.isConnected) {
            this.addConsoleMessage('No hay conexión con el backend', 'error');
            return;
        }

        const code = this.editor.getContent();
        if (!code.trim()) {
            this.addConsoleMessage('El archivo está vacío', 'warning');
            return;
        }

        // UI feedback
        const executeBtn = document.getElementById('executeBtn');
        const originalText = executeBtn.textContent;
        executeBtn.textContent = '⏳ Ejecutando...';
        executeBtn.disabled = true;

        this.addConsoleMessage(`Ejecutando archivo: ${activeFile}`, 'info');
        const startTime = Date.now();

        try {
            const result = await window.apiClient.executeCode(code, activeFile);
            const executionTime = Date.now() - startTime;

            this.displayExecutionResults(result, executionTime);

        } catch (error) {
            console.error('Error executing code:', error);
            this.addConsoleMessage(`Error al ejecutar: ${error.message}`, 'error');
            this.updateStatusMessage('Error de ejecución');
        } finally {
            executeBtn.textContent = originalText;
            executeBtn.disabled = false;
        }
    }

    displayExecutionResults(result, executionTime) {
        // Limpiar errores previos
        this.editor.clearErrors();

        if (result.success) {
            this.addConsoleMessage('✅ Ejecución completada exitosamente', 'success');

            // Mostrar salida del programa
            if (result.output && result.output.length > 0) {
                result.output.forEach(line => {
                    this.addConsoleMessage(line, 'output');
                });
            }

            this.updateStatusMessage('Ejecución exitosa');
        } else {
            this.addConsoleMessage('❌ Ejecución falló con errores', 'error');

            // Marcar errores en el editor
            if (result.errors && result.errors.length > 0) {
                this.editor.markErrors(result.errors);
            }

            this.updateStatusMessage(`${result.errors?.length || 0} errores encontrados`);
        }

        // Actualizar tiempo de ejecución
        document.getElementById('executionTime').textContent = `Tiempo: ${executionTime}ms`;

        // Actualizar reportes
        this.reportsManager.updateReports(result);

        // Cambiar a panel de reportes si hay errores
        if (result.errors && result.errors.length > 0) {
            this.switchPanel('reports');
        }
    }

    // Métodos de gestión de archivos
    createNewFile() {
        this.fileManager.createNewFile();
    }

    openFile() {
        this.fileManager.openFileDialog();
    }

    saveCurrentFile() {
        this.fileManager.saveCurrentFile();
    }

    saveFileAs() {
        this.fileManager.saveFileAs();
    }

    // Gestión de consola
    addConsoleMessage(message, type = 'info') {
        const consoleOutput = document.getElementById('consoleOutput');
        const timestamp = new Date().toLocaleTimeString();

        const messageElement = document.createElement('div');
        messageElement.className = `console-message ${type}`;
        messageElement.innerHTML = `
            <span class="timestamp">[${timestamp}]</span>
            <span class="message">${this.escapeHtml(message)}</span>
        `;

        consoleOutput.appendChild(messageElement);
        consoleOutput.scrollTop = consoleOutput.scrollHeight;

        // Limitar número de mensajes
        const messages = consoleOutput.children;
        if (messages.length > 1000) {
            consoleOutput.removeChild(messages[0]);
        }
    }

    clearConsole() {
        const consoleOutput = document.getElementById('consoleOutput');
        consoleOutput.innerHTML = '';
        this.addConsoleMessage('Consola limpiada', 'info');
    }

    updateStatusMessage(message) {
        document.getElementById('statusMessage').textContent = message;
    }

    escapeHtml(text) {
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }

    // Método para obtener estado del IDE
    getIDEState() {
        return {
            activeFile: this.fileManager.getActiveFile(),
            openFiles: Array.from(this.fileManager.getOpenFiles().keys()),
            hasUnsavedChanges: this.fileManager.hasUnsavedChanges(),
            isConnected: this.isConnected,
            hasErrors: this.reportsManager.hasErrors()
        };
    }
}

// Inicializar IDE cuando el DOM esté listo
document.addEventListener('DOMContentLoaded', () => {
    new IDEController();
});