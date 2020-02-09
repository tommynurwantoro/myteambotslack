class CreateCustomCommands < ActiveRecord::Migration[5.2]
  def change
    create_table :custom_commands, id: false, primary_key: :id do |t|
      t.primary_key :id, :unsigned_integer, auto_increment: true
      t.string :channel_id, null: false
      t.string :command, null: false
      t.string :message, null: false

      t.timestamps null: false
    end
  end
end
