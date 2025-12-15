-- Migração para remover o campo 'imagem' conflitante da tabela academias
-- Execute esta query no seu banco PostgreSQL (Neon)

-- 1. Primeiro, verificar se o campo existe e tem dados
SELECT column_name, data_type, is_nullable
FROM information_schema.columns
WHERE table_name = 'academias' AND column_name = 'imagem';

-- 2. Se o campo existir e estiver vazio ou não for usado, removê-lo
-- ATENÇÃO: Execute apenas se confirmar que o campo não é necessário
ALTER TABLE academias DROP COLUMN IF EXISTS imagem;

-- 3. Verificar a estrutura final da tabela
SELECT column_name, data_type, is_nullable
FROM information_schema.columns
WHERE table_name = 'academias'
ORDER BY ordinal_position;