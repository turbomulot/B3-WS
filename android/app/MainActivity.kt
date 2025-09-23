class MainActivity : AppCompatActivity() {

    lateinit var node: MeshNode
    lateinit var messagesAdapter: ArrayAdapter<String>
    lateinit var peersAdapter: ArrayAdapter<String>
    val messagesList = mutableListOf<String>()
    val peersList = mutableListOf<String>()
    var currentMode = 0

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.layout)

        node = InitNode()

        messagesAdapter = ArrayAdapter(this, android.R.layout.simple_list_item_1, messagesList)
        findViewById<ListView>(R.id.messagesList).adapter = messagesAdapter

        peersAdapter = ArrayAdapter(this, android.R.layout.simple_list_item_1, peersList)
        findViewById<ListView>(R.id.peersList).adapter = peersAdapter

        val spinner = findViewById<Spinner>(R.id.modeDropdown)
        spinner.onItemSelectedListener = object: AdapterView.OnItemSelectedListener {
            override fun onItemSelected(parent: AdapterView<*>, view: View, pos: Int, id: Long) { currentMode = pos }
            override fun onNothingSelected(parent: AdapterView<*>) {}
        }

        findViewById<Button>(R.id.sendButton).setOnClickListener {
            val toID = "peer_id_here"
            val input = findViewById<EditText>(R.id.inputMessage)
            node.Send(toID, input.text.toString().toByteArray(), currentMode)
            messagesList.add("Me: ${input.text}")
            messagesAdapter.notifyDataSetChanged()
            input.text.clear()
        }
    }
}