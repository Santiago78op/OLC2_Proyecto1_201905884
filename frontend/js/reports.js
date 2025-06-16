class ReportsManager {
    constructor() {
        this.currentReports = {
            errors: [],
            symbols: [],
            ast: null
        };
        this.sortOrder = {}; // Para mantener el estado de ordenamiento
        this.astZoom = 1;
        this.astSvg = null;
        this.astData = null;
        this.init();
    }

    init() {
        this.bindEvents();
        this.setupModal();
    }

    bindEvents() {
        // Botón para mostrar reportes
        document.getElementById('showReportsBtn').addEventListener('click', () => {
            this.showReportsModal();
        });

        // Cerrar modal
        document.getElementById('closeReportsModal').addEventListener('click', () => {
            this.hideReportsModal();
        });

        // Cerrar modal al hacer click fuera
        document.getElementById('reportsModal').addEventListener('click', (e) => {
            if (e.target.id === 'reportsModal') {
                this.hideReportsModal();
            }
        });

        // Tabs del modal
        document.querySelectorAll('.modal-tab').forEach(tab => {
            tab.addEventListener('click', () => {
                this.switchModalTab(tab.dataset.tab);
            });
        });

        // Botones de descarga - Errores
        document.getElementById('downloadErrorsBtn').addEventListener('click', () => {
            this.downloadErrors('csv');
        });
        document.getElementById('downloadErrorsExcelBtn').addEventListener('click', () => {
            this.downloadErrors('excel');
        });

        // Botones de descarga - Símbolos
        document.getElementById('downloadSymbolsBtn').addEventListener('click', () => {
            this.downloadSymbols('csv');
        });
        document.getElementById('downloadSymbolsExcelBtn').addEventListener('click', () => {
            this.downloadSymbols('excel');
        });

        // Controles de AST
        document.getElementById('zoomInBtn').addEventListener('click', () => {
            this.zoomAST(1.2);
        });
        document.getElementById('zoomOutBtn').addEventListener('click', () => {
            this.zoomAST(0.8);
        });
        document.getElementById('resetZoomBtn').addEventListener('click', () => {
            this.resetASTZoom();
        });
        document.getElementById('expandAllBtn').addEventListener('click', () => {
            this.expandAllAST();
        });
        document.getElementById('collapseAllBtn').addEventListener('click', () => {
            this.collapseAllAST();
        });

        // Botones de descarga - AST
        document.getElementById('downloadASTSvgBtn').addEventListener('click', () => {
            this.downloadAST('svg');
        });
        document.getElementById('downloadASTPngBtn').addEventListener('click', () => {
            this.downloadAST('png');
        });
        document.getElementById('downloadASTJsonBtn').addEventListener('click', () => {
            this.downloadAST('json');
        });

        // Teclas de escape para cerrar modal
        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape' && document.getElementById('reportsModal').style.display !== 'none') {
                this.hideReportsModal();
            }
        });
    }

    setupModal() {
        // Configuración inicial del modal
        this.hideReportsModal();
    }

    showReportsModal() {
        document.getElementById('reportsModal').style.display = 'flex';
        document.body.style.overflow = 'hidden'; // Prevenir scroll del fondo

        // Refresh data al mostrar
        this.updateAllReports();
    }

    hideReportsModal() {
        document.getElementById('reportsModal').style.display = 'none';
        document.body.style.overflow = 'auto';
    }

    switchModalTab(tabName) {
        // Actualizar tabs
        document.querySelectorAll('.modal-tab').forEach(tab => {
            tab.classList.toggle('active', tab.dataset.tab === tabName);
        });

        // Mostrar/ocultar contenido
        document.querySelectorAll('.modal-tab-content').forEach(content => {
            content.classList.toggle('active', content.id === tabName + 'Tab');
        });

        // Si es AST y hay datos, renderizar
        if (tabName === 'ast' && this.currentReports.ast) {
            setTimeout(() => this.renderAST(), 100);
        }
    }

    updateReports(data) {
        if (!data) return;

        // Mapear los datos del backend a la estructura esperada
        this.currentReports = {
            errors: data.errors || [],
            symbols: data.symbols || data.symbolTable || [],
            ast: data.ast || data.cstSvg || null
        };

        console.log('📊 Reportes actualizados:', {
            errores: this.currentReports.errors.length,
            símbolos: this.currentReports.symbols.length,
            tieneAST: !!this.currentReports.ast
        });

        // Si el modal está abierto, actualizar
        if (document.getElementById('reportsModal').style.display !== 'none') {
            this.updateAllReports();
        }
    }

    updateAllReports() {
        this.updateErrorsTable();
        this.updateSymbolsTable();
        this.updateASTVisualization();
    }

    // ==================== ERRORES ====================
    updateErrorsTable() {
        const tbody = document.getElementById('errorsTableBody');
        const count = document.getElementById('errorsCount');
        const errors = this.currentReports.errors;

        count.textContent = `${errors.length} ${errors.length === 1 ? 'error' : 'errores'}`;

        if (errors.length === 0) {
            tbody.innerHTML = '<tr class="empty-row"><td colspan="5">No hay errores que mostrar</td></tr>';
            return;
        }

        tbody.innerHTML = '';
        errors.forEach((error, index) => {
            const row = document.createElement('tr');
            row.className = 'error-row';

            // Mapear tipos de error a clases CSS
            const typeMapping = {
                'lexical': 'error-type-lexical',
                'syntax': 'error-type-syntax', 
                'semantic': 'error-type-semantic',
                'runtime': 'error-type-runtime'
            };

            const typeClass = typeMapping[error.type] || 'error-type-unknown';
            
            // Nombres en español para mostrar
            const typeNames = {
                'lexical': 'LÉXICO',
                'syntax': 'SINTÁCTICO',
                'semantic': 'SEMÁNTICO', 
                'runtime': 'EJECUCIÓN'
            };

            const typeName = typeNames[error.type] || error.type.toUpperCase();

            row.innerHTML = `
                <td>${index + 1}</td>
                <td class="error-message">${this.escapeHtml(error.message || error.msg || '')}</td>
                <td>${error.line || 0}</td>
                <td>${error.column || 0}</td>
                <td><span class="error-type-cell ${typeClass}">${typeName}</span></td>
            `;

            // Click para ir a la línea
            row.addEventListener('click', () => {
                this.goToLocation(error.line, error.column);
                this.hideReportsModal();
            });

            tbody.appendChild(row);
        });
    }

    downloadErrors(format) {
        const errors = this.currentReports.errors;
        if (errors.length === 0) {
            alert('No hay errores para descargar');
            return;
        }

        const data = errors.map((error, index) => ({
            'No.': index + 1,
            'Descripción': error.message || error.description || '',
            'Línea': error.line || 0,
            'Columna': error.column || 0,
            'Tipo': (error.type || 'ERROR').toUpperCase()
        }));

        if (format === 'csv') {
            this.downloadCSV(data, 'errores.csv');
        } else if (format === 'excel') {
            this.downloadExcel(data, 'errores.xlsx', 'Errores');
        }
    }

    // ==================== SÍMBOLOS ====================
    updateSymbolsTable() {
        const tbody = document.getElementById('symbolsTableBody');
        const count = document.getElementById('symbolsCount');
        const symbols = this.currentReports.symbols;

        count.textContent = `${symbols.length} ${symbols.length === 1 ? 'símbolo' : 'símbolos'}`;

        if (symbols.length === 0) {
            tbody.innerHTML = '<tr class="empty-row"><td colspan="5">No hay símbolos que mostrar</td></tr>';
            return;
        }

        tbody.innerHTML = '';
        symbols.forEach((symbol, index) => {
            const row = document.createElement('tr');
            row.className = 'symbol-row';

            // Determinar tipo y clase CSS
            const symbolType = this.getSymbolTypeClass(symbol.type);
            const scopeClass = this.getScopeClass(symbol.scope);

            row.innerHTML = `
                <td><span class="symbol-name">${symbol.name || `SYM_${index + 1}`}</span></td>
                <td><span class="symbol-type-cell ${symbolType}">${this.getSymbolTypeDisplay(symbol.type)}</span></td>
                <td>${symbol.type || 'unknown'}</td>
                <td><span class="symbol-scope ${scopeClass}">${symbol.scope || 'global'}</span></td>
                <td><span class="symbol-location">${symbol.line || 0}:${symbol.column || 0}</span></td>
            `;

            // Click para ir a la línea
            if (symbol.line > 0) {
                row.addEventListener('click', () => {
                    this.goToLocation(symbol.line, symbol.column);
                    this.hideReportsModal();
                });
            }

            tbody.appendChild(row);
        });
    }

    // Función para determinar la clase CSS del tipo
    getSymbolTypeClass(type) {
        if (!type) return 'symbol-type-unknown';
        
        if (type.includes('Embebida')) return 'symbol-type-builtin';
        if (type === 'variable' || type === 'int' || type === 'string' || type === 'bool' || type === 'float') return 'symbol-type-variable';
        if (type === 'function' || type.includes('function')) return 'symbol-type-function';
        if (type === 'struct') return 'symbol-type-struct';
        
        return 'symbol-type-variable';
    }

    // Función para determinar la clase CSS del scope
    getScopeClass(scope) {
        if (!scope || scope === 'global') return 'scope-global';
        if (scope.includes('func') || scope.includes('function')) return 'scope-function';
        if (scope.includes('struct')) return 'scope-struct';
        return 'scope-local';
    }

    // Función para mostrar el tipo de símbolo
    getSymbolTypeDisplay(type) {
        if (!type) return 'DESCONOCIDO';
        
        if (type.includes('Embebida')) return 'INCORPORADA';
        if (type === 'variable' || type === 'int' || type === 'string' || type === 'bool' || type === 'float') return 'VARIABLE';
        if (type === 'function' || type.includes('function')) return 'FUNCIÓN';
        if (type === 'struct') return 'ESTRUCTURA';
        
        return type.toUpperCase();
    }

    downloadSymbols(format) {
        const symbols = this.currentReports.symbols;
        if (symbols.length === 0) {
            alert('No hay símbolos para descargar');
            return;
        }

        const data = symbols.map((symbol, index) => ({
            'ID': symbol.id || symbol.name || `SYM_${index + 1}`,
            'Tipo de Símbolo': (symbol.type || 'VARIABLE').toUpperCase(),
            'Tipo de Dato': symbol.dataType || symbol.valueType || 'unknown',
            'Ámbito': symbol.scope || symbol.ambito || 'global',
            'Línea': symbol.line || 0,
            'Columna': symbol.column || 0
        }));

        if (format === 'csv') {
            this.downloadCSV(data, 'tabla_simbolos.csv');
        } else if (format === 'excel') {
            this.downloadExcel(data, 'tabla_simbolos.xlsx', 'Símbolos');
        }
    }

    // ==================== AST ====================
    updateASTVisualization() {
        const container = document.getElementById('astVisualization');
        const ast = this.currentReports.ast || this.currentReports.cstSvg;

        console.log('🌳 Actualizando AST:', {
            tieneAST: !!ast,
            tipoAST: typeof ast,
            longitudAST: ast ? ast.length : 0
        });

        if (!ast) {
            container.innerHTML = `
                <div class="empty-ast">
                    <div class="empty-ast-icon">🌳</div>
                    <p>No hay AST que mostrar</p>
                    <small>Ejecuta código para generar el árbol de sintaxis</small>
                </div>
            `;
            return;
        }

        this.astData = ast;

        // Si la tab de AST está activa, renderizar inmediatamente
        const astTab = document.getElementById('astTab');
        if (astTab && astTab.classList.contains('active')) {
            this.renderAST();
        }
    }

    renderAST() {
        if (!this.astData) {
            console.log('❌ No hay datos de AST para renderizar');
            return;
        }

        const container = document.getElementById('astVisualization');
        container.innerHTML = '';

        console.log('🎨 Renderizando AST...');

        // Verificar si es SVG
        if (typeof this.astData === 'string' && this.astData.includes('<svg')) {
            this.renderSVGAST(container);
        } else if (typeof this.astData === 'object') {
            this.renderJSONAST(container);
        } else {
            this.renderTextAST(container);
        }
    }

    // Método para renderizar AST en formato SVG
    renderSVGAST(container) {
        console.log('📊 Renderizando SVG AST');
        
        const wrapper = document.createElement('div');
        wrapper.className = 'ast-svg-wrapper';
        wrapper.style.cssText = `
            width: 100%;
            height: 100%;
            display: flex;
            justify-content: center;
            align-items: center;
            background: #1e1e1e;
            overflow: auto;
            padding: 20px;
            box-sizing: border-box;
        `;

        // Insertar el SVG directamente
        wrapper.innerHTML = this.astData;

        // Ajustar el SVG para que sea responsive
        const svg = wrapper.querySelector('svg');
        if (svg) {
            svg.style.cssText = `
                max-width: 100%;
                max-height: 100%;
                width: auto;
                height: auto;
                border: 1px solid #3e3e42;
                border-radius: 8px;
                background: #1e1e1e;
            `;
            
            // Agregar zoom con scroll
            this.addSVGZoomControls(wrapper, svg);
        }

        container.appendChild(wrapper);
        this.astZoom = 1;
    }

    // Método para renderizar AST en formato JSON/Objeto
    renderJSONAST(container) {
        console.log('🔧 Renderizando JSON AST');
        
        const wrapper = document.createElement('div');
        wrapper.className = 'ast-json-wrapper';
        wrapper.style.cssText = `
            width: 100%;
            height: 100%;
            overflow: auto;
            padding: 20px;
            background: #1e1e1e;
            font-family: 'Consolas', monospace;
        `;

        // Crear visualización en árbol
        const treeHTML = this.createTreeVisualization(this.astData);
        wrapper.innerHTML = treeHTML;

        container.appendChild(wrapper);
    }

    // Método para renderizar AST en formato texto
    renderTextAST(container) {
        console.log('📝 Renderizando texto AST');
        
        const wrapper = document.createElement('div');
        wrapper.className = 'ast-text-wrapper';
        wrapper.style.cssText = `
            width: 100%;
            height: 100%;
            overflow: auto;
            padding: 20px;
            background: #1e1e1e;
            font-family: 'Consolas', monospace;
            color: #d4d4d4;
            white-space: pre-wrap;
        `;

        wrapper.textContent = this.astData;
        container.appendChild(wrapper);
    }

    // Agregar controles de zoom para SVG
    addSVGZoomControls(wrapper, svg) {
        let scale = 1;
        const minScale = 0.1;
        const maxScale = 3;

        wrapper.addEventListener('wheel', (e) => {
            e.preventDefault();
            
            const delta = e.deltaY > 0 ? 0.9 : 1.1;
            scale = Math.max(minScale, Math.min(maxScale, scale * delta));
            
            svg.style.transform = `scale(${scale})`;
            this.astZoom = scale;
        });
    }

    // Crear visualización en árbol para JSON
    createTreeVisualization(data, level = 0) {
        const indent = '  '.repeat(level);
        let html = '';

        if (typeof data === 'object' && data !== null) {
            if (Array.isArray(data)) {
                html += `<div class="ast-array" style="margin-left: ${level * 20}px;">[\n`;
                data.forEach((item, index) => {
                    html += this.createTreeVisualization(item, level + 1);
                    if (index < data.length - 1) html += ',';
                    html += '\n';
                });
                html += `${indent}]</div>`;
            } else {
                html += `<div class="ast-object" style="margin-left: ${level * 20}px;">{\n`;
                Object.keys(data).forEach((key, index, keys) => {
                    html += `<div class="ast-property" style="margin-left: ${(level + 1) * 20}px;">`;
                    html += `<span class="ast-key" style="color: #9cdcfe;">"${key}"</span>: `;
                    html += this.createTreeVisualization(data[key], level + 1);
                    if (index < keys.length - 1) html += ',';
                    html += '</div>\n';
                });
                html += `${indent}}</div>`;
            }
        } else {
            const color = typeof data === 'string' ? '#ce9178' : 
                        typeof data === 'number' ? '#b5cea8' : 
                        typeof data === 'boolean' ? '#569cd6' : '#d4d4d4';
            
            const value = typeof data === 'string' ? `"${data}"` : String(data);
            html += `<span style="color: ${color};">${value}</span>`;
        }

        return html;
    }

    zoomAST(factor) {
        const container = document.getElementById('astVisualization');
        const svg = container.querySelector('svg');
        
        if (svg) {
            this.astZoom *= factor;
            this.astZoom = Math.max(0.1, Math.min(3, this.astZoom));
            svg.style.transform = `scale(${this.astZoom})`;
        }
    }

    resetASTZoom() {
        const container = document.getElementById('astVisualization');
        const svg = container.querySelector('svg');
        
        if (svg) {
            this.astZoom = 1;
            svg.style.transform = 'scale(1)';
        }
    }

    expandAllAST() {
        // Por simplicidad, solo actualizamos el mensaje
        console.log('Expandir todos los nodos');
    }

    collapseAllAST() {
        // Por simplicidad, solo actualizamos el mensaje
        console.log('Colapsar todos los nodos');
    }

    downloadAST(format) {
        if (!this.astData) {
            alert('No hay AST para descargar');
            return;
        }

        switch (format) {
            case 'svg':
                this.downloadASTSVG();
                break;
            case 'png':
                this.downloadASTPNG();
                break;
            case 'json':
                this.downloadASTJSON();
                break;
        }
    }

    downloadASTSVG() {
        if (!this.astData || !this.astData.includes('<svg')) {
            alert('No hay SVG AST para descargar');
            return;
        }

        const blob = new Blob([this.astData], { type: 'image/svg+xml' });
        this.downloadBlob(blob, 'ast.svg');
    }

    downloadASTJSON() {
        if (!this.astData) {
            alert('No hay AST para descargar');
            return;
        }

        let jsonContent;
        if (typeof this.astData === 'string') {
            jsonContent = JSON.stringify({ ast: this.astData }, null, 2);
        } else {
            jsonContent = JSON.stringify(this.astData, null, 2);
        }

        const blob = new Blob([jsonContent], { type: 'application/json' });
        this.downloadBlob(blob, 'ast.json');
    }


    downloadASTPNG() {
        const container = document.getElementById('astVisualization');
        const svg = container.querySelector('svg');
        
        if (!svg) {
            alert('No hay SVG para convertir a PNG');
            return;
        }

        // Crear canvas y convertir SVG a PNG
        const canvas = document.createElement('canvas');
        const ctx = canvas.getContext('2d');
        const data = new XMLSerializer().serializeToString(svg);
        
        const img = new Image();
        img.onload = () => {
            canvas.width = img.width || 800;
            canvas.height = img.height || 600;
            
            // Fondo negro
            ctx.fillStyle = '#1e1e1e';
            ctx.fillRect(0, 0, canvas.width, canvas.height);
            
            ctx.drawImage(img, 0, 0);
            
            canvas.toBlob((blob) => {
                this.downloadBlob(blob, 'ast.png');
            });
        };
        
        img.src = 'data:image/svg+xml;base64,' + btoa(unescape(encodeURIComponent(data)));
    }

    // ==================== UTILIDADES ====================
    downloadCSV(data, filename) {
        if (!data.length) return;

        const headers = Object.keys(data[0]);
        const csvContent = [
            headers.join(','),
            ...data.map(row =>
                headers.map(header => {
                    const value = row[header];
                    const stringValue = String(value || '');
                    // Escapar comillas y envolver en comillas si contiene comas
                    return stringValue.includes(',') || stringValue.includes('"')
                        ? `"${stringValue.replace(/"/g, '""')}"`
                        : stringValue;
                }).join(',')
            )
        ].join('\n');

        const blob = new Blob(['\ufeff' + csvContent], { type: 'text/csv;charset=utf-8;' });
        this.downloadBlob(blob, filename);
    }

    downloadExcel(data, filename, sheetName) {
        if (!data.length) return;

        const wb = XLSX.utils.book_new();
        const ws = XLSX.utils.json_to_sheet(data);

        // Ajustar ancho de columnas
        const cols = Object.keys(data[0]).map(key => ({ wch: 20 }));
        ws['!cols'] = cols;

        XLSX.utils.book_append_sheet(wb, ws, sheetName);
        XLSX.writeFile(wb, filename);
    }

    downloadBlob(blob, filename) {
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.style.display = 'none';
        a.href = url;
        a.download = filename;
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);
        URL.revokeObjectURL(url);
    }

    goToLocation(line, column = 1) {
        if (window.ideController?.editor) {
            window.ideController.editor.goToLine(line, column);
        }
    }

    escapeHtml(text) {
        if (typeof text !== 'string') {
            text = String(text);
        }
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }

    clearReports() {
        this.currentReports = {
            errors: [],
            symbols: [],
            ast: null
        };

        if (document.getElementById('reportsModal').style.display !== 'none') {
            this.updateAllReports();
        }
    }

    hasErrors() {
        return this.currentReports.errors.length > 0;
    }

    getErrorCount() {
        return this.currentReports.errors.length;
    }
}

// Función global para ordenar tablas
function sortTable(tableId, columnIndex) {
    const table = document.getElementById(tableId);
    const tbody = table.querySelector('tbody');
    const rows = Array.from(tbody.querySelectorAll('tr:not(.empty-row)'));

    if (rows.length === 0) return;

    const sortKey = `${tableId}_${columnIndex}`;
    const currentOrder = window.reportsManager?.sortOrder?.[sortKey] || 'asc';
    const newOrder = currentOrder === 'asc' ? 'desc' : 'asc';

    if (window.reportsManager) {
        window.reportsManager.sortOrder[sortKey] = newOrder;
    }

    rows.sort((a, b) => {
        const aValue = a.children[columnIndex]?.textContent?.trim() || '';
        const bValue = b.children[columnIndex]?.textContent?.trim() || '';

        // Intentar convertir a número si es posible
        const aNum = parseFloat(aValue);
        const bNum = parseFloat(bValue);

        let comparison = 0;
        if (!isNaN(aNum) && !isNaN(bNum)) {
            comparison = aNum - bNum;
        } else {
            comparison = aValue.localeCompare(bValue);
        }

        return newOrder === 'asc' ? comparison : -comparison;
    });

    // Actualizar iconos de ordenamiento
    const headers = table.querySelectorAll('th');
    headers.forEach((th, index) => {
        const icon = th.querySelector('.sort-icon');
        if (icon) {
            if (index === columnIndex) {
                icon.textContent = newOrder === 'asc' ? '↑' : '↓';
            } else {
                icon.textContent = '↕️';
            }
        }
    });

    // Reordenar filas
    rows.forEach(row => tbody.appendChild(row));
}