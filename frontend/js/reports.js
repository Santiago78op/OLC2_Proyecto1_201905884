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

        this.currentReports = {
            errors: data.errors || [],
            symbols: data.symbols || data.symbolTable || [],
            ast: data.ast || null
        };

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

            const typeClass = `error-type-${error.type || 'error'}`;

            row.innerHTML = `
                <td>${index + 1}</td>
                <td>${this.escapeHtml(error.message || error.description || '')}</td>
                <td>${error.line || 0}</td>
                <td>${error.column || 0}</td>
                <td><span class="error-type-cell ${typeClass}">${(error.type || 'ERROR').toUpperCase()}</span></td>
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
            tbody.innerHTML = '<tr class="empty-row"><td colspan="6">No hay símbolos que mostrar</td></tr>';
            return;
        }

        tbody.innerHTML = '';
        symbols.forEach((symbol, index) => {
            const row = document.createElement('tr');
            row.className = 'symbol-row';

            const typeClass = `symbol-type-${symbol.type || 'variable'}`;

            row.innerHTML = `
                <td>${symbol.id || symbol.name || `SYM_${index + 1}`}</td>
                <td><span class="symbol-type-cell ${typeClass}">${(symbol.type || 'VARIABLE').toUpperCase()}</span></td>
                <td>${symbol.dataType || symbol.valueType || 'unknown'}</td>
                <td>${symbol.scope || symbol.ambito || 'global'}</td>
                <td>${symbol.line || 0}</td>
                <td>${symbol.column || 0}</td>
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
        const ast = this.currentReports.ast;

        if (!ast) {
            container.innerHTML = `
                <div class="empty-ast">
                    <p>No hay AST que mostrar</p>
                    <small>Ejecuta código para generar el árbol de sintaxis</small>
                </div>
            `;
            return;
        }

        this.astData = ast;

        // Si la tab de AST está activa, renderizar inmediatamente
        const astTab = document.getElementById('astTab');
        if (astTab.classList.contains('active')) {
            this.renderAST();
        }
    }

    renderAST() {
        if (!this.astData) return;

        const container = document.getElementById('astVisualization');
        container.innerHTML = '';

        const width = container.clientWidth || 800;
        const height = container.clientHeight || 600;

        // Crear SVG
        this.astSvg = d3.select(container)
            .append('svg')
            .attr('class', 'ast-svg')
            .attr('width', width)
            .attr('height', height);

        // Crear grupo principal con zoom
        const g = this.astSvg.append('g');

        // Configurar zoom
        const zoom = d3.zoom()
            .scaleExtent([0.1, 3])
            .on('zoom', (event) => {
                g.attr('transform', event.transform);
            });

        this.astSvg.call(zoom);

        // Crear layout de árbol
        const tree = d3.tree()
            .size([width - 100, height - 100])
            .separation((a, b) => (a.parent === b.parent ? 1 : 2) / a.depth);

        // Convertir datos
        const root = d3.hierarchy(this.astData);
        tree(root);

        // Crear enlaces
        const links = g.selectAll('.link')
            .data(root.links())
            .enter().append('path')
            .attr('class', 'link')
            .attr('d', d3.linkHorizontal()
                .x(d => d.y + 50)
                .y(d => d.x + 50));

        // Crear nodos
        const nodes = g.selectAll('.node')
            .data(root.descendants())
            .enter().append('g')
            .attr('class', 'node')
            .attr('transform', d => `translate(${d.y + 50},${d.x + 50})`);

        // Círculos de nodos
        nodes.append('circle')
            .attr('r', 20)
            .on('click', (event, d) => {
                if (d.data.line > 0) {
                    this.goToLocation(d.data.line, d.data.column);
                    this.hideReportsModal();
                }
            });

        // Texto de nodos
        nodes.append('text')
            .text(d => d.data.type || 'Node')
            .style('font-size', '10px');

        // Tooltips
        nodes.append('title')
            .text(d => {
                const info = [];
                if (d.data.type) info.push(`Tipo: ${d.data.type}`);
                if (d.data.value) info.push(`Valor: ${d.data.value}`);
                if (d.data.line) info.push(`Línea: ${d.data.line}`);
                if (d.data.column) info.push(`Columna: ${d.data.column}`);
                return info.join('\n');
            });

        this.astZoom = 1;
    }

    zoomAST(factor) {
        if (!this.astSvg) return;

        this.astZoom *= factor;
        this.astSvg.transition()
            .duration(300)
            .call(d3.zoom().scaleBy, factor);
    }

    resetASTZoom() {
        if (!this.astSvg) return;

        this.astZoom = 1;
        this.astSvg.transition()
            .duration(500)
            .call(d3.zoom().transform, d3.zoomIdentity);
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
        if (!this.astSvg) return;

        const svgElement = this.astSvg.node();
        const serializer = new XMLSerializer();
        const svgString = serializer.serializeToString(svgElement);

        const blob = new Blob([svgString], { type: 'image/svg+xml' });
        this.downloadBlob(blob, 'ast.svg');
    }

    downloadASTPNG() {
        if (!this.astSvg) return;

        const svgElement = this.astSvg.node();
        const canvas = document.createElement('canvas');
        const ctx = canvas.getContext('2d');

        const data = (new XMLSerializer()).serializeToString(svgElement);
        const img = new Image();

        img.onload = () => {
            canvas.width = img.width;
            canvas.height = img.height;
            ctx.fillStyle = '#1e1e1e';
            ctx.fillRect(0, 0, canvas.width, canvas.height);
            ctx.drawImage(img, 0, 0);

            canvas.toBlob((blob) => {
                this.downloadBlob(blob, 'ast.png');
            });
        };

        img.src = 'data:image/svg+xml;base64,' + btoa(data);
    }

    downloadASTJSON() {
        const jsonString = JSON.stringify(this.astData, null, 2);
        const blob = new Blob([jsonString], { type: 'application/json' });
        this.downloadBlob(blob, 'ast.json');
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