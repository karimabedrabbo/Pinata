//
//  CameraViewController.swift
//  Vibin
//
//  Created by Karim Abedrabbo on 9/1/18.
//  Copyright © 2018 Karim Abedrabbo. All rights reserved.
//


import UIKit
import AVFoundation
import Photos
import SCRecorder
import AssetsLibrary
import MobileCoreServices

@objc protocol CameraViewControllerProtocol: NSObjectProtocol {
    @objc optional func recorder(_ recorder: CameraViewController?, didFinishPickingMediaWith videoUrl: URL?)

    @objc optional func recorderDidCancel(_ recorder: CameraViewController?)
}

let AlbumTitle = "PINATA"

class CameraViewController: UIViewController {
    
    // MARK: - properties
    
    weak var delegate: CameraViewControllerProtocol?
    
    internal var previewView: UIView!
    internal var focusView: FocusIndicatorView?
    internal var recorder: SCRecorder!
    
    
    internal var hasOrientaionLocked = false
    internal var setToDelete = false
    

    internal var longPressStartPoint: CGPoint = .zero
    internal var startZoom: CGFloat = 0.0
    
    internal var minDuration: CGFloat = 0.5
    internal var maxDuration: CGFloat = 25.0
    
    var recordViewBoundsObserver: NSKeyValueObservation?
    var blurredRecordViewBoundsObserver: NSKeyValueObservation?
    
    
    var progressViewBars: [UIView] = []
    var colorRotation: [UIColor] = [.green, .blue, .red,  .magenta, .yellow, .cyan, .brown, .purple, .orange]
    
    var alphaIsIncreasing: Bool = false
    var readyToAddBar: Bool = false
    var lastCompletionRatio: CGFloat = 0.0
    var completionRatios: [CGFloat] = []
    var activeBarWidthConstraint: NSLayoutConstraint?
    
    var innerShapeLayer: CAShapeLayer =  {
        let innerShape = CAShapeLayer()
        innerShape.backgroundColor = UIColor.clear.cgColor
        
        //innerShape.fillColor = UIColor.clear.cgColor
        return innerShape
    }()
    
    var outerShapeLayer: CAShapeLayer =  {
        let outerShape = CAShapeLayer()
        outerShape.lineWidth = 4.0
        outerShape.backgroundColor = UIColor.clear.cgColor
        outerShape.strokeColor = UIColor.white.cgColor
        outerShape.fillColor = UIColor.clear.cgColor
        return outerShape
    }()
    
    
    
    var flipButton: UIButton = {
        let flip = UIButton(type: .custom)
        flip.setImage(UIImage(named: "flip_button"), for: .normal)
        flip.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)

        //flip.sizeToFit()
        flip.addTarget(self, action: #selector(handleFlipButton(_:)), for: .touchUpInside)
        flip.isHidden = false
        flip.translatesAutoresizingMaskIntoConstraints = false
        return flip
    }()
    
    var flashButton: UIButton = {
        let flash = UIButton(type: .custom)
        flash.setImage(UIImage(named: "flash_off_button"), for: .normal)
        flash.setImage(UIImage(named: "flash_on_button"), for: .selected)
        flash.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)
        //flash.sizeToFit()
        flash.addTarget(self, action: #selector(handleFlashModeButton(_:)), for: .touchUpInside)
        flash.isHidden = false
        flash.translatesAutoresizingMaskIntoConstraints = false
        return flash
    }()
    
    var mediaButton: UIButton = {
        let media = UIButton(type: .custom)
        media.setImage(UIImage(named: "media_button"), for: .normal)
        media.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)

        //media.sizeToFit()
        media.addTarget(self, action: #selector(handleMediaButton(_:)), for: .touchUpInside)
        media.isHidden = false
        media.translatesAutoresizingMaskIntoConstraints = false
        return media
    }()
    
    var textButton: UIButton = {
        let text = UIButton(type: .custom)
        text.setImage(UIImage(named: "text_button_thin"), for: .normal)
        text.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)
        
        //media.sizeToFit()
        text.addTarget(self, action: #selector(handleTextButton(_:)), for: .touchUpInside)
        text.isHidden = false
        text.translatesAutoresizingMaskIntoConstraints = false
        return text
    }()
    
    //doesnt work as a view but works as a button WITH NO TARGET TAP HANDLED BY GESTURES because there is also a long hold gesture
    var deleteSegmentView: UIButton = {
//        let deleteSegment = UIImageView(image: UIImage(named: "delete_segment_button"))
        let deleteSegment = UIButton(type: .custom)
        deleteSegment.setImage( UIImage(named: "delete_segment_button"), for: .normal)
        //deleteSegment.sizeToFit()
        deleteSegment.contentMode = .center
        deleteSegment.isHidden = true
        deleteSegment.translatesAutoresizingMaskIntoConstraints = false
        
        deleteSegment.layer.cornerRadius = 30
        return deleteSegment
        
    }()
    
    var cancelButton: UIButton = {
        let cancel = UIButton(type: .custom)
        cancel.setImage(UIImage(named: "cancel_button"), for: .normal)
        cancel.imageEdgeInsets = UIEdgeInsets(top: 5, left: 5, bottom: 5, right: 5)

        //cancel.sizeToFit()
        cancel.addTarget(self, action: #selector(handleCancelButton(_:)), for: .touchUpInside)
        cancel.isHidden = false
        cancel.translatesAutoresizingMaskIntoConstraints = false
        return cancel
    }()
    
    
    var doneRecordingButton: UIButton = {
        let doneRecording = UIButton(type: .custom)
        var completeImage = UIImage(named: "complete_button")
        //completeImage = completeImage?.resizeImage(newHeight: 15.0)
        doneRecording.setImage(completeImage, for: .normal)
        //doneRecording.sizeToFit()
        doneRecording.setTitle("  Next", for: .normal)
        doneRecording.titleLabel?.font = doneRecording.titleLabel?.font.withSize(20)
        doneRecording.isEnabled = false
        doneRecording.addTarget(self, action: #selector(handleDoneRecordingButton(_:)), for: .touchUpInside)
        doneRecording.backgroundColor = UIColor(red: 210.0/255.0, green: 40/255.0, blue: 40.0/255.0, alpha: 1.0)
        doneRecording.titleLabel?.font = UIFont.boldSystemFont(ofSize: 16.0)
        
        doneRecording.layer.cornerRadius = 30.0
        doneRecording.isHidden = true
        doneRecording.translatesAutoresizingMaskIntoConstraints = false
        return doneRecording
    }()
    
    
    var progressView: UIView = {
        let progress = UIView()
        progress.backgroundColor = UIColor.white.withAlphaComponent(0.3)
        progress.contentMode = .left
        progress.layer.cornerRadius = 15.0
        progress.translatesAutoresizingMaskIntoConstraints = false
        progress.clipsToBounds = true
        return progress
    }()
    
    var minimumTimeView: UIView = {
        let minimumTime = UIView()
        minimumTime.backgroundColor = UIColor.white.withAlphaComponent(0.6)
        minimumTime.translatesAutoresizingMaskIntoConstraints = false
        return minimumTime
    }()
    
   
    var blurredRecordView: UIView = {
        let blurEffect = UIBlurEffect(style: UIBlurEffect.Style.regular)
        
        let blurView = UIVisualEffectView(effect: blurEffect)
        
        
        blurView.backgroundColor = UIColor.clear
        blurView.translatesAutoresizingMaskIntoConstraints = false
        return blurView
    }()
    
    var recordView: UIView = {
        
        let recordView = UIView()
        recordView.backgroundColor = UIColor.clear
        recordView.translatesAutoresizingMaskIntoConstraints = false
  
        return recordView
    }()
    
    var gestureView: UIView = {
        
        // gestures
        let gesture = UIView()
        gesture.translatesAutoresizingMaskIntoConstraints = false
        gesture.backgroundColor = .clear

        return gesture
        
    }()
    
    deinit {
        print("recordViewDeinitialized")
    }
    
    // MARK: - view lifecycle
    
    func setupRecorder() {
        
        self.recorder = SCRecorder()
        recorder.initializeSessionLazily = false
        prepareSession()
        //recorder.captureSessionPreset = SCRecorderTools.bestCaptureSessionPresetCompatibleWithAllDevices()
        
        recorder.captureSessionPreset = SCRecorderTools.bestCaptureSessionPreset(for: self.recorder.device, withMaxSize: CGSize(width: 1920, height: 1080))
        
        if self.maxDuration != 0 {
            recorder.maxRecordDuration = CMTime(seconds: Double(self.maxDuration), preferredTimescale: CMTimeScale(1.0))
        }
        
        
        recorder.flashMode = .off
        //recorder.device = AVCaptureDevice.Position.front
        //recorder.videoConfiguration.size = CGSizeMake(800,800)
        recorder.autoSetVideoOrientation = false
        recorder.delegate = self
        recorder.videoOrientation = AVCaptureVideoOrientation.portrait
        
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.navigationItem.backBarButtonItem = UIBarButtonItem(title: "", style: .plain, target: nil, action: nil)
        self.navigationItem.backBarButtonItem?.tintColor = UIColor.white
        
        let screenBounds = UIScreen.main.bounds
        
        self.view.backgroundColor = UIColor.black
        self.view.autoresizingMask = [.flexibleWidth, .flexibleHeight, .flexibleLeftMargin, .flexibleRightMargin, .flexibleTopMargin, .flexibleBottomMargin]
        
        self.setupRecorder()
        
        //set hidden features
        //flashButton.isHidden = !recorder.deviceHasFlash
        
        
        // preview
        self.previewView = UIView()

        self.previewView.translatesAutoresizingMaskIntoConstraints = false
        self.previewView.backgroundColor = UIColor.clear
        
        
        
        recorder.previewView = self.previewView
        //recorder.previewView?.frame = self.previewView.bounds
        self.previewView.layer.addSublayer(recorder.previewLayer)
        self.view.addSubview(previewView)
      
        
        do {
            try recorder.prepare()
        }
        catch {
            print("Prepare error: \(error.localizedDescription )")
        }
        
        
        self.recordViewBoundsObserver = self.recordView.observe(\.bounds, changeHandler: { [weak self] (view, change) in
            self?.resetRecordViewPaths()
        })
        
        self.blurredRecordViewBoundsObserver = self.blurredRecordView.observe(\.bounds, changeHandler: { [weak self] (view, change) in
            self?.resetBlurredRecordViewPaths()
        })
        
        
       
        
        self.previewView.addSubview(progressView)
        
       
        self.previewView.addSubview(gestureView)
        
        self.progressView.addSubview(minimumTimeView)
        
        self.previewView.addSubview(mediaButton)
        self.previewView.addSubview(textButton)
        self.previewView.addSubview(flipButton)
        self.previewView.addSubview(deleteSegmentView)
        self.previewView.addSubview(doneRecordingButton)
        self.previewView.addSubview(cancelButton)
        self.previewView.addSubview(flashButton)
        
        
        //note that record button is attached to the view itself. When it is attached to the preview view the .setNeedsLayoutConstraint just stops working
        self.view.addSubview(recordView)
        recordView.layer.addSublayer(outerShapeLayer)
        self.recordView.addSubview(blurredRecordView)
        blurredRecordView.layer.mask = innerShapeLayer
        
        
        
        
        

        NSLayoutConstraint.activate([
      
        self.previewView.topAnchor.constraint(equalTo: self.view.topAnchor),
        self.previewView.bottomAnchor.constraint(equalTo: self.view.bottomAnchor),
        self.previewView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor),
        self.previewView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor),
        self.previewView.heightAnchor.constraint(equalTo: self.view.heightAnchor),
        self.previewView.widthAnchor.constraint(equalTo: self.view.widthAnchor),
        
        self.gestureView.topAnchor.constraint(equalTo: self.previewView.topAnchor),
        self.gestureView.bottomAnchor.constraint(equalTo: self.previewView.bottomAnchor),
        self.gestureView.leadingAnchor.constraint(equalTo: self.previewView.leadingAnchor),
        self.gestureView.trailingAnchor.constraint(equalTo: self.previewView.trailingAnchor),
        self.gestureView.heightAnchor.constraint(equalTo: self.previewView.heightAnchor),
        self.gestureView.widthAnchor.constraint(equalTo: self.previewView.widthAnchor),
            
        self.progressView.topAnchor.constraint(equalTo: self.previewView.topAnchor, constant: 30),
        self.progressView.widthAnchor.constraint(equalTo: self.previewView.widthAnchor, constant: -30),
        self.progressView.heightAnchor.constraint(equalToConstant: 30),
        self.progressView.centerXAnchor.constraint(equalTo: self.previewView.centerXAnchor),
        
        
        self.minimumTimeView.topAnchor.constraint(equalTo: self.progressView.topAnchor),
//        self.minimumTimeView.bottomAnchor.constraint(equalTo: self.progressView.bottomAnchor),
//
            self.minimumTimeView.leadingAnchor.constraint(equalTo: self.progressView.leadingAnchor, constant: (screenBounds.width - 30.0) * (CGFloat(self.minDuration / self.maxDuration))),
        
        self.minimumTimeView.widthAnchor.constraint(equalToConstant: 2.0),
        
        self.minimumTimeView.heightAnchor.constraint(equalTo: self.progressView.heightAnchor),
        
        self.cancelButton.widthAnchor.constraint(equalToConstant: 50),
        self.cancelButton.heightAnchor.constraint(equalToConstant: 50),
        self.cancelButton.leadingAnchor.constraint(equalTo: self.previewView.leadingAnchor, constant: 25),
        self.cancelButton.topAnchor.constraint(equalTo: self.progressView.bottomAnchor, constant: 40),
        
        self.flashButton.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
        self.flashButton.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
        self.flashButton.trailingAnchor.constraint(equalTo: self.previewView.trailingAnchor, constant: -25),
        self.flashButton.topAnchor.constraint(equalTo: self.cancelButton.topAnchor),
        
        self.flipButton.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
        self.flipButton.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
        self.flipButton.trailingAnchor.constraint(equalTo: self.flashButton.trailingAnchor),
        self.flipButton.topAnchor.constraint(equalTo: self.flashButton.bottomAnchor, constant: 40),
        
        self.mediaButton.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
        self.mediaButton.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
        self.mediaButton.trailingAnchor.constraint(equalTo: self.flashButton.trailingAnchor),
        self.mediaButton.topAnchor.constraint(equalTo: self.flipButton.bottomAnchor, constant: 40),
        
        self.textButton.widthAnchor.constraint(equalTo: self.cancelButton.widthAnchor),
        self.textButton.heightAnchor.constraint(equalTo: self.cancelButton.heightAnchor),
        self.textButton.trailingAnchor.constraint(equalTo: self.flashButton.trailingAnchor),
        self.textButton.topAnchor.constraint(equalTo: self.mediaButton.bottomAnchor, constant: 40),
        

        
      
        
        self.recordView.centerXAnchor.constraint(equalTo: self.view.centerXAnchor),
        self.recordView.bottomAnchor.constraint(equalTo: self.view.bottomAnchor, constant: -30),
        self.recordView.heightAnchor.constraint(equalToConstant: 130),
        self.recordView.widthAnchor.constraint(equalToConstant: 130),
        
        self.blurredRecordView.centerXAnchor.constraint(equalTo: self.recordView.centerXAnchor),
        self.blurredRecordView.centerYAnchor.constraint(equalTo: self.recordView.centerYAnchor),
        self.blurredRecordView.heightAnchor.constraint(equalTo: self.recordView.heightAnchor, constant: 0),
        self.blurredRecordView.widthAnchor.constraint(equalTo: self.recordView.widthAnchor, constant: 0),
        
        self.deleteSegmentView.widthAnchor.constraint(equalToConstant: 60),
        self.deleteSegmentView.heightAnchor.constraint(equalToConstant: 60),
        self.deleteSegmentView.trailingAnchor.constraint(equalTo: self.recordView.leadingAnchor, constant: -25),
        self.deleteSegmentView.centerYAnchor.constraint(equalTo: self.recordView.centerYAnchor),
        
        self.doneRecordingButton.widthAnchor.constraint(equalToConstant: 90),
        self.doneRecordingButton.heightAnchor.constraint(equalToConstant: 60),
        self.doneRecordingButton.leadingAnchor.constraint(equalTo: self.recordView.trailingAnchor, constant: 15),
        self.doneRecordingButton.centerYAnchor.constraint(equalTo: self.recordView.centerYAnchor),
        
        ])
        
        
        

        
        let recordLongPressRecognizer = UILongPressGestureRecognizer(target: self, action: #selector(handleRecordLongPressRecognizer(_:)))
        recordLongPressRecognizer.delegate = self
        recordLongPressRecognizer.minimumPressDuration = 0.01
        recordLongPressRecognizer.allowableMovement = 50.0
        self.recordView.addGestureRecognizer(recordLongPressRecognizer)
        
        
        let recordTapRecognizer = UITapGestureRecognizer(target: self, action: #selector(handleRecordTapRecognizer(_:)))
        recordTapRecognizer.delegate = self
        recordTapRecognizer.numberOfTapsRequired = 1
        recordTapRecognizer.numberOfTouchesRequired = 1
        recordView.addGestureRecognizer(recordTapRecognizer)
        
        
        self.enableDisableTapGesture()
        
        
        let deleteTapPressGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(handleDeleteSegmentTap(_:)))
        deleteTapPressGestureRecognizer.delegate = self
        deleteTapPressGestureRecognizer.numberOfTapsRequired = 1
        deleteTapPressGestureRecognizer.numberOfTouchesRequired = 1
    deleteSegmentView.addGestureRecognizer(deleteTapPressGestureRecognizer)
        
        let deleteLongPressGestureRecognizer = UILongPressGestureRecognizer(target: self, action: #selector(handleDeleteLongPressGestureRecognizer(_:)))
        deleteLongPressGestureRecognizer.delegate = self
        deleteLongPressGestureRecognizer.minimumPressDuration = 0.4
        deleteSegmentView.addGestureRecognizer(deleteLongPressGestureRecognizer)
        
        
        
        let focusTapGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(handleFocusTapGestureRecognizer(_:)))
        focusTapGestureRecognizer.delegate = self
        focusTapGestureRecognizer.numberOfTapsRequired = 1
        gestureView.addGestureRecognizer(focusTapGestureRecognizer)
        
        
        let panGestureRecognizer = UIPanGestureRecognizer(target: self, action: #selector(handlePanGestureRecognizer(_:)))
        panGestureRecognizer.delegate = self
        gestureView.addGestureRecognizer(panGestureRecognizer)
       

        
    }
    
   
    
    override func viewDidLayoutSubviews() {
        super.viewDidLayoutSubviews()
        if recorder != nil {
            recorder.previewViewFrameChanged()
        }
    }
    
    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)
        
        
        //prepareSession()
            
        
        recorder.startRunning()

        
    }
    
    
    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)
        self.navigationController?.setNavigationBarHidden(true, animated: animated)
    }
    
    override func viewWillDisappear(_ animated: Bool) {
        super.viewWillDisappear(animated)
        self.navigationController?.setNavigationBarHidden(false, animated: animated)

        self.recorder.stopRunning()
        //self.executeResetAction()

    }
    
    override func viewDidDisappear(_ animated: Bool) {
        super.viewDidDisappear(animated)
        
        
        if let delegate = self.delegate, let cancelFunction = delegate.recorderDidCancel {
            cancelFunction(self)
        }
        
        
    }
    
    override var prefersStatusBarHidden: Bool {
        return true
    }
    
    func shouldAutorotate() -> Bool {
        return false
    }
    
    override var supportedInterfaceOrientations: UIInterfaceOrientationMask {
        return [.portrait, .portraitUpsideDown]
    }
    
    
    override var preferredInterfaceOrientationForPresentation: UIInterfaceOrientation {
        return .portrait
    }
}

//MARK: - Custom Functions
extension CameraViewController {
    
    internal func startCapture() {
        if self.recorder.ratioRecorded < 1.0 {
            hideAllButtons()
            if !hasOrientaionLocked {
                recorder.autoSetVideoOrientation = false
                hasOrientaionLocked = true
                let currentOrientation: UIInterfaceOrientation = UIApplication.shared.statusBarOrientation
                UIDevice.current.setValue(currentOrientation.rawValue, forKey: "Orientation")
            }
            if !recorder.isRecording {
                recorder.record()
            }
        }
    }
    
    internal func hideAllButtons() {
        self.cancelButton.isHidden = true
        self.flashButton.isHidden = true
        self.flipButton.isHidden = true
        self.mediaButton.isHidden = true
        self.textButton.isHidden = true
        self.deleteSegmentView.isHidden = true
        self.doneRecordingButton.isHidden = true
    }
    
    internal func showAllButtons() {
        self.cancelButton.isHidden = false
        self.flashButton.isHidden = false
        self.flipButton.isHidden = false
        self.mediaButton.isHidden = false
        self.textButton.isHidden = false
        self.deleteSegmentView.isHidden = false
        if doneRecordingButton.isEnabled {
            self.doneRecordingButton.isHidden = false
        } else {
            self.doneRecordingButton.isHidden = true
        }
    }
    
    func scaleCircleAnimatedSingle(newRadius: CGFloat, shape : CAShapeLayer, duration: Double) {
        
        let newPath: UIBezierPath = UIBezierPath(arcCenter: CGPoint(x: newRadius * 2, y: newRadius * 2), radius: newRadius, startAngle: 0.0, endAngle: 2 * CGFloat(Double.pi), clockwise: false)
        
        
        let newBounds : CGRect = CGRect(x: 0, y: 0, width: 2 * newRadius, height: 2 * newRadius)
        let pathAnim : CABasicAnimation = CABasicAnimation(keyPath: "path")
        
        pathAnim.toValue = newPath.cgPath
        
        let boundsAnim : CABasicAnimation = CABasicAnimation(keyPath: "bounds")
        boundsAnim.toValue = NSValue(cgRect: newBounds)
        
        let anims: CAAnimationGroup = CAAnimationGroup()
        anims.animations = [pathAnim, boundsAnim]
        anims.isRemovedOnCompletion = false
        anims.duration = duration
        anims.fillMode = CAMediaTimingFillMode.both
        anims.autoreverses = false
        anims.repeatCount = 1
        shape.add(anims, forKey: nil)
        
    }
    
    func scaleCircleAnimatedRepeat(newRadius: CGFloat, shape : CAShapeLayer, duration: Double){
        
        let newPath: UIBezierPath = UIBezierPath(arcCenter: CGPoint(x: newRadius * 2, y: newRadius * 2), radius: newRadius, startAngle: 0.0, endAngle: 2 * CGFloat(Double.pi), clockwise: false)
        
        
        let newBounds : CGRect = CGRect(x: 0, y: 0, width: 2 * newRadius, height: 2 * newRadius)
        let pathAnim : CABasicAnimation = CABasicAnimation(keyPath: "path")
        
        pathAnim.toValue = newPath.cgPath
        
        let boundsAnim : CABasicAnimation = CABasicAnimation(keyPath: "bounds")
        boundsAnim.toValue = NSValue(cgRect: newBounds)
        
        let anims: CAAnimationGroup = CAAnimationGroup()
        anims.animations = [pathAnim, boundsAnim]
        anims.isRemovedOnCompletion = true
        anims.duration = 0.3
        anims.fillMode = CAMediaTimingFillMode.forwards
        anims.autoreverses = true
        anims.repeatCount = Float(Int.max)
        shape.add(anims, forKey: nil)
    }
    
    internal func animateRecordingButton() {
        self.scaleCircleAnimatedSingle(newRadius: recordView.bounds.width / 2.1, shape: self.outerShapeLayer, duration: 0.1)

        self.scaleCircleAnimatedRepeat(newRadius: recordView.bounds.width / 1.9, shape: self.outerShapeLayer, duration: 3.0)
        
    }
    
   
    
    internal func pauseCapture(completion: (() -> ())?) {
        showAllButtons()
        outerShapeLayer.strokeColor = UIColor.white.cgColor
        innerShapeLayer.isHidden = false
        self.outerShapeLayer.removeAllAnimations()
        
        recordView.setNeedsUpdateConstraints()
        if recorder.isRecording {
            recorder.pause { [weak self] in
                self?.lastCompletionRatio = self?.recorder.ratioRecorded ?? 0.0
                if self?.lastCompletionRatio ?? CGFloat(0) > CGFloat(0) {
                    self?.completionRatios.append((self?.lastCompletionRatio)!)
                }
                
                self?.enableDisableTapGesture()
                
                if let done = completion {
                    done()
                }
            }
        } else {
            if let done = completion {
                done()
            }
        }
    }
    
    internal func endCapture() {

        if CGFloat(self.recorder.session?.duration.seconds ?? 0.0) > self.minDuration {
            self.pauseCapture { [weak self] in
                //guard let self = self else { return }
                
                self?.readyToAddBar = true
                self?.activeBarWidthConstraint = nil
                
                
                //SCRecordSessionManager.sharedInstance()?.saveRecord(self?.recorder.session)
                //self?.session = self?.recorder.session
                
                let newNav = UINavigationController(rootViewController: VideoPlayerViewController(asset: (self?.recorder.session?.assetRepresentingSegments())!, outputUrl: (self?.recorder.session?.outputUrl)!, completionRatios: self?.completionRatios ?? [], colorRotation: self?.colorRotation ?? [], style: "cameraSegmentsVideo"))
                
                self?.present(newNav, animated: true, completion: nil)
        
                
                }
            
            }
        
    }
    
    internal func prepareSession() {
        if self.recorder.session == nil {
            let session = SCRecordSession()
            session.fileType = AVFileType.mov.rawValue
            self.recorder.session = session
        }
    }
    
    func updateProgressView() {
        let currentCompletionRatio = recorder.ratioRecorded
        if self.readyToAddBar {
            let colorIndex: Int = self.progressViewBars.count
            let newBar = UIView()
            newBar.translatesAutoresizingMaskIntoConstraints = false
            let newColor = self.colorRotation[colorIndex % self.colorRotation.count]
            newBar.backgroundColor = newColor
            deleteSegmentView.backgroundColor = newColor
            outerShapeLayer.strokeColor = newColor.cgColor
            innerShapeLayer.isHidden = true
            animateRecordingButton()
            self.progressView.addSubview(newBar)
            
            //self.activeBarWidthConstraint?.isActive = false
            self.activeBarWidthConstraint = newBar.widthAnchor.constraint(equalToConstant: 0)
            if !self.progressViewBars.isEmpty {
                NSLayoutConstraint.activate([
                    newBar.topAnchor.constraint(equalTo: self.progressView.topAnchor),
                    newBar.bottomAnchor.constraint(equalTo: self.progressView.bottomAnchor),
                    newBar.heightAnchor.constraint(equalTo: self.progressView.heightAnchor),
                    newBar.leadingAnchor.constraint(equalTo: self.progressViewBars.last!.trailingAnchor),
                    self.activeBarWidthConstraint!
                    ])
            } else {
                NSLayoutConstraint.activate([
                newBar.topAnchor.constraint(equalTo: self.progressView.topAnchor),
                newBar.bottomAnchor.constraint(equalTo: self.progressView.bottomAnchor),
                newBar.heightAnchor.constraint(equalTo: self.progressView.heightAnchor),
                newBar.leadingAnchor.constraint(equalTo: self.progressView.leadingAnchor),
                self.activeBarWidthConstraint!
                ])
            }
            
            self.progressViewBars.append(newBar)

            self.readyToAddBar = false
            
        }
        
        if !self.progressViewBars.isEmpty {
          let newWidth = (currentCompletionRatio - self.lastCompletionRatio) * self.progressView.frame.width
            if !newWidth.isNaN && newWidth > 0 {
                self.activeBarWidthConstraint?.constant = newWidth
            } else if newWidth < 0 {
                executeSegmentDeletionProgressUpdate()
            } else {
                self.activeBarWidthConstraint?.constant = 0.0
            }
            
            self.progressView.setNeedsLayout()
            
        }
        if self.alphaIsIncreasing {
            self.progressView.alpha += 0.03
        } else {
            self.progressView.alpha -= 0.03
        }
        if self.progressView.alpha >= 1.0 {
            self.alphaIsIncreasing = false
        } else if self.progressView.alpha <= 0.7 {
            self.alphaIsIncreasing = true
        }
    }

    func checkMaxSegmentDuration() {
        if self.maxDuration != 0 {
            if let timeElapsed = recorder.session?.duration {
                //end capture if the duration of the session greater than the max_duration
                if CGFloat(timeElapsed.seconds) > self.maxDuration {
                    self.endCapture()
                }
                
                //enable done button if duration is greater than minimum
                if !self.doneRecordingButton.isEnabled, CGFloat(timeElapsed.seconds) > self.minDuration {
                    self.doneRecordingButton.isEnabled = true
                    //self.doneRecordingButton.isHidden = false
                }
            }
        }
    }
}

//MARK: - Handle Button Functions
extension CameraViewController {
    @objc internal func handleFlipButton(_ button: UIButton) {
        if self.recorder.device == .front {
             recorder.captureSessionPreset = SCRecorderTools.bestCaptureSessionPreset(for: .back, withMaxSize: CGSize(width: 1920, height: 1080))
        } else {
            recorder.captureSessionPreset = SCRecorderTools.bestCaptureSessionPreset(for: .front, withMaxSize: CGSize(width: 1920, height: 1080))
            
        }
       
        recorder.switchCaptureDevices()

        self.flashButton.isSelected = false
    }
    
    internal func executeResetAction() {
        self.doneRecordingButton.isEnabled = false
        self.doneRecordingButton.isHidden = true
        for bar in self.progressViewBars {
            bar.removeFromSuperview()
        }
        self.progressViewBars = []
        self.completionRatios = []
        self.readyToAddBar = true
        self.deleteSegmentView.isHidden = true
        self.lastCompletionRatio = 0.0
        self.activeBarWidthConstraint = nil
        self.outerShapeLayer.strokeColor = UIColor.white.cgColor
        innerShapeLayer.isHidden = false
        self.outerShapeLayer.removeAllAnimations()
        
        self.enableDisableTapGesture()
        
//
//        if self.recorder.session != nil {
//
//
//            // If the thisRecordSession was saved, we don't want to completely destroy it
//            if SCRecordSessionManager.sharedInstance() != nil {
//                if SCRecordSessionManager.sharedInstance()!.isSaved(self.recorder.session) {
//                    self.recorder.session?.endSegment(withInfo: nil, completionHandler: nil)
//                } else {
//                    self.recorder.session?.cancel(nil)
//                }
//            }
//
//            self.recorder.session = nil
//        }
    }
    
    
    @objc internal func handleCancelButton(_ button: UIButton) {
        executeResetAction()
        self.recorder.unprepare()
        self.dismiss(animated: true, completion: nil)
    }
    
    
    func handleDeleteAllAction() {
        executeResetAction()
        prepareSession()
    }
    
    
    @objc internal func handleFlashModeButton(_ button: UIButton) {
        if recorder.deviceHasFlash{
        if button.isSelected {
            recorder.flashMode = SCFlashMode.off
            button.isSelected = false
        } else {
            recorder.flashMode = SCFlashMode.light
            button.isSelected = true
            }
        }
    }
    
    @objc internal func handleDeleteSegmentTap(_ gestureRecognizer: UIGestureRecognizer) {
        self.doneRecordingButton.isEnabled = false
        self.doneRecordingButton.isHidden = true
        
        //hackery needed because of mismatch between completing recording pausing if recorder is recording
        self.setToDelete = true
        if !self.recorder.isRecording && self.setToDelete {
            self.executeSegmentDeletion()
        } else {
            //will consequently complete segment where setToDelete will be evaluated again
            self.pauseCapture(completion: nil)
        }
    }
    
    
    internal func enableDisableTapGesture() {
        if self.lastCompletionRatio > 0 {
            for gesture in self.recordView.gestureRecognizers ?? [] {
                if gesture is UITapGestureRecognizer {
                    gesture.isEnabled = false
                } else if gesture is UILongPressGestureRecognizer {
                    let longPress = gesture as! UILongPressGestureRecognizer
                    longPress.minimumPressDuration = 0.01
                }
            }
        } else {
            for gesture in self.recordView.gestureRecognizers ?? [] {
                if gesture is UITapGestureRecognizer {
                    gesture.isEnabled = true
                } else if gesture is UILongPressGestureRecognizer {
                    let longPress = gesture as! UILongPressGestureRecognizer
                    longPress.minimumPressDuration = 0.35
                }
            }
          self.colorRotation = colorRotation.shuffled()
        }
    }
    
    internal func executeSegmentDeletionProgressUpdate() {
        //pop segment time and remove from marker from superview
        if let lastBar = self.progressViewBars.popLast() {
            lastBar.removeFromSuperview()
            
            //tidy up things for updating progress view
            self.activeBarWidthConstraint = nil
            self.readyToAddBar = true
            self.completionRatios.popLast()
            self.lastCompletionRatio = self.completionRatios.last ?? 0.0
            self.outerShapeLayer.strokeColor = UIColor.white.cgColor
            innerShapeLayer.isHidden = false
            self.outerShapeLayer.removeAllAnimations()
            
            self.enableDisableTapGesture()
            
            if self.progressViewBars.isEmpty {
                self.deleteSegmentView.isHidden = true
            } else {
                self.deleteSegmentView.backgroundColor = self.progressViewBars.last!.backgroundColor
            }
            
        }
    }
    
    internal func executeSegmentDeletion() {
        if let session = self.recorder.session {

            executeSegmentDeletionProgressUpdate()
            
            session.removeLastSegment()
            if CGFloat(session.duration.seconds) < self.minDuration {
                self.doneRecordingButton.isEnabled = false
                self.doneRecordingButton.isHidden = true
            } else {
                self.doneRecordingButton.isEnabled = true
                self.doneRecordingButton.isHidden = false
            }
            self.setToDelete = false
            print("deleted segment")
        }
    }
    
    @objc internal func handleTextButton(_ button: UIButton) {
        textButton.isEnabled = false
//        self.navigationController?.pushViewController(FinishUploadPostViewController(style: "textOnly"), animated: true)
        self.textButton.isEnabled = true
        
    }
    
    @objc internal func handleMediaButton(_ button: UIButton) {
        mediaButton.isEnabled = false
        if let mediaBrowser = self.initializeMediaBrowser() {
        
            present(mediaBrowser, animated: true) { [weak self] in
                print("completed media presenter")
                self?.mediaButton.isEnabled = true
            }
        }
    }
    
    @objc internal func handleDoneRecordingButton(_ button: UIButton) {
        self.endCapture()
    }
}

// MARK: - recorder delegate functions

extension CameraViewController: SCRecorderDelegate {
    
    func recorder(_ recorder: SCRecorder, didAppendVideoSampleBufferIn session: SCRecordSession) {
        updateProgressView()
        checkMaxSegmentDuration()
        print("appeded video sample")
    }
    
    func recorder(_ recorder: SCRecorder, didSkipVideoSampleBufferIn session: SCRecordSession) {
        print("Skipped video buffer")
    }
    
    func recorderWillStartFocus(_ recorder: SCRecorder) {
        focusView?.startAnimation()
    }

    func recorderDidEndFocus(_ recorder: SCRecorder) {
        focusView?.stopAnimation()
    }
    
    func recorder(_ recorder: SCRecorder, didSkipAudioSampleBufferIn session: SCRecordSession) {
        print("Skipped audio buffer")
    }
    
    func recorder(_ recorder: SCRecorder, didReconfigureVideoInput videoInputError: Error?) {
        if let anError = videoInputError {
            print("Reconfigured video input: \(anError)")
        }
    }
    
    func recorder(_ recorder: SCRecorder, didReconfigureAudioInput audioInputError: Error?) {
        if let anError = audioInputError {
            print("Reconfigured audio input: \(anError)")
        }
    }
    
    
    
    func recorder(_ recorder: SCRecorder, didInitializeAudioIn session: SCRecordSession, error: Error?) {
        if error == nil {
            print("Initialized audio in record session")
        } else {
            print("Failed to initialize audio in record session: \(error?.localizedDescription ?? "")")
        }
    }
    
    func recorder(_ recorder: SCRecorder, didInitializeVideoIn session: SCRecordSession, error: Error?) {
        if error == nil {
            print("Initialized video in record session")
        } else {
            print("Failed to initialize video in record session: \(error?.localizedDescription ?? "")")
        }
    }
    
    
    func recorder(_ recorder: SCRecorder, didBeginSegmentIn session: SCRecordSession, error: Error?) {
        if let anError = error {
            print("could not begin segment with error: \(anError)")
        } else {
            self.readyToAddBar = true
        }
    }
    
    func recorder(_ recorder: SCRecorder, didComplete segment: SCRecordSessionSegment?, in session: SCRecordSession, error: Error?) {
        if let anError = error {
            print("error completing segment with error: \(anError))")
        } else {
            print("completed segment")
            if self.setToDelete {
                self.executeSegmentDeletion()
            }
            self.alphaIsIncreasing = false
            self.progressView.alpha = 1.0
            
            //self.updateProgressView()
            //ready to add after progressView updated
            self.readyToAddBar = true
        }
    }
}

//MARK: - ImagePicker
extension CameraViewController: UIImagePickerControllerDelegate, UINavigationControllerDelegate, CameraViewControllerProtocol {
    func initializeMediaBrowser() -> UIImagePickerController? {
        //Validations
        if UIImagePickerController.isSourceTypeAvailable(.photoLibrary) == false {
            return nil
        }
        
        let mediaUI = UIImagePickerController()
        mediaUI.sourceType = .photoLibrary
        mediaUI.mediaTypes = [kUTTypeMovie as String, kUTTypeImage as String]
//        mediaUI.mediaTypes = [kUTTypeMovie as String, kUTTypeJPEG as String, kUTTypePNG as String, kUTTypePic]
        //mediaUI.videoMaximumDuration = TimeInterval(self.maxDuration - (self.maxDuration * recorder.ratioRecorded))
        mediaUI.videoMaximumDuration = TimeInterval(self.maxDuration)
        mediaUI.allowsEditing = true
        
        mediaUI.delegate = self
        
        
        return mediaUI
    }
    
    
    
    func recorderDidCancel(_ recorder: CameraViewController?) {
        print("recorder did cancel")
    }
    
    func imagePickerController(_ picker: UIImagePickerController, didFinishPickingMediaWithInfo info: [UIImagePickerController.InfoKey : Any]) {
            if let pickedImage = info[UIImagePickerController.InfoKey.originalImage] as? UIImage {
                picker.dismiss(animated: true)
                self.present(UINavigationController(rootViewController: VideoPlayerViewController(image: pickedImage, style: "photoLibraryImage")), animated: true)
            } else if let url = info[UIImagePickerController.InfoKey.mediaType] as? URL {
                picker.dismiss(animated: true)
    
                    let segment = SCRecordSessionSegment(url: url, info: nil)
                    if CGFloat(segment.duration.seconds) > self.minDuration {
    
                        self.present(UINavigationController(rootViewController: VideoPlayerViewController(asset: segment.asset!, outputUrl: url, completionRatios: [0.0, 1.0], colorRotation: self.colorRotation, style: "photoLibraryMedia")), animated: true)
                    } else {
                        let alert = UIAlertController(title: "Video Too Short", message: "The video you uploaded was too short, please edit your video in the picker to be shorter than \(Int(self.minDuration.rounded())) seconds", preferredStyle: .actionSheet)
                        alert.addAction(UIAlertAction(title: "Got it", style: .default, handler: nil))
                        self.present(alert, animated: true)
                    }
    
            } else {
                picker.dismiss(animated: true)
            }
    
        }
    

}



// MARK: - UIGestureRecognizerDelegate

extension CameraViewController: UIGestureRecognizerDelegate {
    
    //tap
    @objc internal func handleRecordTapRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
        gestureRecognizer.isEnabled = false
        self.recorder.capturePhoto { [weak self] (error, image) in
            if error != nil {
                print("Could not capture photo with error: \(error)")
            } else if image == nil {
                print("Image captured is nil")
            } else {
                self?.present(UINavigationController(rootViewController: VideoPlayerViewController(image: image!, style: "cameraPhoto")), animated: true, completion: nil)

            }
        }
    }
    
    //long press
    @objc internal func handleRecordLongPressRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
            switch gestureRecognizer.state {
            case .began:
                self.startCapture()
                self.longPressStartPoint = gestureRecognizer.location(in: self.previewView)
                self.startZoom = CGFloat(self.recorder.videoZoomFactor)
                break
            case .changed:
                let newGesturePoint = gestureRecognizer.location(in: self.previewView)
                
                recordView.center = newGesturePoint
                
                
                let deltaY = self.longPressStartPoint.y / newGesturePoint.y
                var zoomScale: CGFloat = 1.0
                if deltaY > 1.0 {
                    zoomScale = pow(deltaY, 1.5)
                } else {
                    zoomScale = pow(deltaY, 10.0)
                }
                let newZoom = (zoomScale * self.startZoom)
                self.recorder.videoZoomFactor = min(max(1.0, CGFloat(newZoom)), CGFloat(self.recorder.maxVideoZoomFactor / 4))
                break
            case .ended:
                fallthrough
            case .cancelled:
                fallthrough
            case .failed:
                self.pauseCapture(completion: nil)
                
                fallthrough
            default:
                break
            }
    }
    
    
         @objc internal func handleDeleteLongPressGestureRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
    
    
           let alert = UIAlertController(title: "Delete All?", message: "", preferredStyle: .actionSheet)
    
           alert.addAction(UIAlertAction(title: "Keep", style: .default, handler: nil))
           alert.addAction(UIAlertAction(title: "Delete All", style: .destructive) { [weak self] (alertAction) in
                 self?.handleDeleteAllAction()
             })
    
    
           self.present(alert, animated: true)
    
         }
    
    //tap
    @objc internal func handleFocusTapGestureRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
        if !recorder.isAdjustingFocus{
        let tapPoint = gestureRecognizer.location(in: self.previewView)
        let focusSize = CGSize(width: 30, height: 30)
        let focusRect = CGRect(x: tapPoint.x - focusSize.width/2, y: tapPoint.y - focusSize.height/2, width: focusSize.width, height: focusSize.height)
        self.focusView = FocusIndicatorView(frame: focusRect)
        self.view.addSubview(focusView!)
        
        let recorderPoint = recorder.convertToPointOfInterest(fromViewCoordinates: tapPoint)
        recorder.autoFocus(at: recorderPoint)
        }
    }
    
    @objc internal func handlePanGestureRecognizer(_ gestureRecognizer: UIGestureRecognizer) {
        switch gestureRecognizer.state {
        case .began:
            self.longPressStartPoint = gestureRecognizer.location(in: self.previewView)
            self.startZoom = CGFloat(self.recorder.videoZoomFactor)
            break
        case .changed:
            let newGesturePoint = gestureRecognizer.location(in: self.previewView)
            
            
            
            let deltaY = self.longPressStartPoint.y / newGesturePoint.y
            var zoomScale: CGFloat = 1.0
            if deltaY > 1.0 {
                zoomScale = pow(deltaY, 1.5)
            } else {
                zoomScale = pow(deltaY, 3.5)
            }
            let newZoom = (zoomScale * self.startZoom)
            self.recorder.videoZoomFactor = min(max(1.0, CGFloat(newZoom)), CGFloat(self.recorder.maxVideoZoomFactor / 4))
            break
        case .ended:
            fallthrough
        case .cancelled:
            fallthrough
        case .failed:
            fallthrough
        default:
            break
        }
    }
    
    
}

//recorder button
extension CameraViewController {
    func resetRecordViewPaths() {
        print(view.bounds.width)
        print(recordView.bounds.width)
        print(recordView.center)
        let recordViewCenter = CGPoint(x: recordView.bounds.width / 2, y: recordView.bounds.height / 2)
        let outerPath = UIBezierPath(arcCenter: recordViewCenter, radius: recordView.bounds.width / 3, startAngle: 0.0, endAngle: 2.0 * .pi, clockwise: false)
    
        
        outerShapeLayer.path = outerPath.cgPath
        
    }
    
    func resetBlurredRecordViewPaths() {
        
        print(blurredRecordView.bounds.width)
        print(blurredRecordView.center)
        let blurredRecordViewCenter = CGPoint(x: blurredRecordView.bounds.width / 2, y: blurredRecordView.bounds.height / 2)
        let innerPath = UIBezierPath(arcCenter: blurredRecordViewCenter, radius: blurredRecordView.bounds.width / 4, startAngle: 0.0, endAngle: 2.0 * .pi, clockwise: false)
        innerShapeLayer.path = innerPath.cgPath
    }
}
